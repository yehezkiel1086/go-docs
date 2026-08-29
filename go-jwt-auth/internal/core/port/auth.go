package port

import "context"

type AuthService interface {
	Login(ctx context.Context, email, password string) (accessToken string, refreshToken string, err error)
	RefreshToken(ctx context.Context, refreshToken string) (accessToken string, newRefreshToken string, err error)
	Logout(ctx context.Context, accessToken string) error
}