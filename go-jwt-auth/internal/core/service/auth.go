package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/yehezkiel1086/go-jwt-auth/internal/adapter/config"
	"github.com/yehezkiel1086/go-jwt-auth/internal/adapter/storage/redis"
	"github.com/yehezkiel1086/go-jwt-auth/internal/core/port"
	"github.com/yehezkiel1086/go-jwt-auth/internal/core/util"
)

type AuthService struct {
	conf     *config.JWT
	cache    *redis.Redis
	userRepo port.UserRepository
}

func NewAuthService(conf *config.JWT, cache *redis.Redis, userRepo port.UserRepository) *AuthService {
	return &AuthService{conf: conf, cache: cache, userRepo: userRepo}
}

func (s *AuthService) refreshTokenCacheKey(userID uint) string {
	return fmt.Sprintf("refresh_token:%d", userID)
}

func (s *AuthService) storeRefreshToken(ctx context.Context, userID uint, token string, ttl time.Duration) error {
	return s.cache.Set(ctx, s.refreshTokenCacheKey(userID), []byte(token), ttl)
}

func (s *AuthService) revokeRefreshToken(ctx context.Context, userID uint) error {
	return s.cache.Del(ctx, s.refreshTokenCacheKey(userID))
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, string, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	if err := util.ComparePassword(user.Password, password); err != nil {
		return "", "", errors.New("invalid credentials")
	}

	accessToken, err := util.GenerateToken("access", user, s.conf)
	if err != nil {
		return "", "", err
	}

	refreshToken, ttl, err := util.GenerateTokenWithTTL("refresh", user, s.conf)
	if err != nil {
		return "", "", err
	}

	if err := s.storeRefreshToken(ctx, user.ID, refreshToken, ttl); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (string, string, error) {
	claims, err := util.ParseToken(refreshToken, s.conf.RefreshTokenSecret)
	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	cacheKey := s.refreshTokenCacheKey(claims.UserID)
	stored, err := s.cache.Get(ctx, cacheKey)
	if err != nil {
		return "", "", errors.New("refresh token expired or not found")
	}

	// reject if token doesn't match — possible reuse attack
	if string(stored) != refreshToken {
		// revoke stored token as a precaution
		s.revokeRefreshToken(ctx, claims.UserID)
		return "", "", errors.New("refresh token reuse detected")
	}

	user, err := s.userRepo.GetUserByID(ctx, claims.UserID)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	// revoke old refresh token before issuing new one
	if err := s.revokeRefreshToken(ctx, claims.UserID); err != nil {
		return "", "", err
	}

	accessToken, err := util.GenerateToken("access", user, s.conf)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, ttl, err := util.GenerateTokenWithTTL("refresh", user, s.conf)
	if err != nil {
		return "", "", err
	}

	if err := s.storeRefreshToken(ctx, user.ID, newRefreshToken, ttl); err != nil {
		return "", "", err
	}

	return accessToken, newRefreshToken, nil
}

func (s *AuthService) Logout(ctx context.Context, accessToken string) error {
	claims, err := util.ParseToken(accessToken, s.conf.AccessTokenSecret)
	if err != nil {
		return errors.New("invalid access token")
	}
	return s.revokeRefreshToken(ctx, claims.UserID)
}