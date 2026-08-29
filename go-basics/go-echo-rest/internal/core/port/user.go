package port

import (
	"context"

	"github.com/yehezkiel1086/go-docs/go-echo-rest/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetByID(ctx context.Context, id int64) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) (domain.User, error)
	DeleteUser(ctx context.Context, id int64) error
	HardDeleteUser(ctx context.Context, id int64) error
	ListUsers(ctx context.Context, offset, limit int) ([]domain.User, error)
	CountUsers(ctx context.Context, status string) (int64, error)
}

type UserService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUserByID(ctx context.Context, id int64) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) (domain.User, error)
	DeleteUser(ctx context.Context, id int64) error
	HardDeleteUser(ctx context.Context, id int64) error
	ListUsers(ctx context.Context, offset, limit int) ([]domain.User, error)
	CountUsers(ctx context.Context, status string) (int64, error)
}
