package port

import (
	"context"

	"github.com/yehezkiel1086/go-socket-chat/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.CreateUserRes, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

type UserService interface {
	Register(ctx context.Context, user *domain.User) (*domain.CreateUserRes, error)
	Login(ctx context.Context, creds *domain.LoginUserReq) (*domain.LoginUserRes, error)
}
