package service

import (
	"context"
	"encoding/json"
	"fmt"
	"go-gin-next-oauth/internal/model"
	"go-gin-next-oauth/internal/repository"
	"go-gin-next-oauth/internal/util"
	"io"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

type AuthService struct {
	oauthConfig *oauth2.Config
	repo        *repository.UserRepository
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		oauthConfig: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		},
		repo: repository.NewUserRepository(db),
	}
}

func (s *AuthService) GetGoogleLoginURL() string {
	return s.oauthConfig.AuthCodeURL("randomstate", oauth2.AccessTypeOffline)
}

func (s *AuthService) HandleGoogleCallback(ctx context.Context, code string) (map[string]interface{}, string, error) {
	token, err := s.oauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, "", fmt.Errorf("failed to exchange token: %w", err)
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var userInfo map[string]interface{}
	json.Unmarshal(body, &userInfo)

	// Persist or update user
	user := &model.User{
		Email:    userInfo["email"].(string),
		Name:     userInfo["name"].(string),
		Picture:  userInfo["picture"].(string),
		Provider: "google",
	}
	s.repo.CreateOrUpdate(user)

	// Generate JWT
	jwtToken, _ := util.GenerateJWT(user.Email)

	return userInfo, jwtToken, nil
}
