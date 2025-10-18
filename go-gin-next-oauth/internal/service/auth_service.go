package service

import (
	"context"
	"encoding/json"
	"fmt"
	"go-gin-next-oauth/internal/util"
	"io"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthService struct {
	oauthConfig *oauth2.Config
}

func NewAuthService() *AuthService {
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

	jwtToken, _ := util.GenerateJWT(userInfo["email"].(string))

	return userInfo, jwtToken, nil
}
