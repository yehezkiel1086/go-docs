package repository

import (
	"context"
	"database/sql"

	"github.com/yehezkiel1086/go-docs/go-echo-rest/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-docs/go-echo-rest/internal/adapter/storage/postgres/gen"
	"github.com/yehezkiel1086/go-docs/go-echo-rest/internal/core/domain"
)

type UserRepository struct {
	db *postgres.DB
	q  *gen.Queries
}

func NewUserRepository(db *postgres.DB, q *gen.Queries) *UserRepository {
	return &UserRepository{
		db: db,
		q:  q,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	params := gen.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     sql.NullString{String: user.Role, Valid: user.Role != ""},
		Status:   sql.NullString{String: user.Status, Valid: user.Status != ""},
	}

	created, err := r.q.CreateUser(ctx, params)
	if err != nil {
		return domain.User{}, err
	}
	return toDomainUser(created), nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (domain.User, error) {
	user, err := r.q.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return toDomainUser(user), nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := r.q.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return toDomainUser(user), nil
}

func (r *UserRepository) ListUsers(ctx context.Context, status *string, limit, offset int32) ([]domain.User, error) {
	users, err := r.q.ListUsers(ctx, gen.ListUsersParams{
		Status: sql.NullString{String: func() string {
			if status != nil {
				return *status
			}
			return ""
		}(), Valid: status != nil},
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	result := make([]domain.User, len(users))
	for i, u := range users {
		result[i] = toDomainUser(u)
	}
	return result, nil
}

func (r *UserRepository) CountUsers(ctx context.Context, status *string) (int64, error) {
	return r.q.CountUsers(ctx, sql.NullString{
		String: func() string {
			if status != nil {
				return *status
			}
			return ""
		}(),
		Valid: status != nil,
	})
}

func (r *UserRepository) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	params := gen.UpdateUserParams{
		ID:       int64(user.ID),
		Name:     sql.NullString{String: user.Name, Valid: user.Name != ""},
		Email:    sql.NullString{String: user.Email, Valid: user.Email != ""},
		Password: sql.NullString{String: user.Password, Valid: user.Password != ""},
		Role:     sql.NullString{String: user.Role, Valid: user.Role != ""},
		Status:   sql.NullString{String: user.Status, Valid: user.Status != ""},
	}

	updated, err := r.q.UpdateUser(ctx, params)
	if err != nil {
		return domain.User{}, err
	}
	return toDomainUser(updated), nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	return r.q.DeleteUser(ctx, id)
}

func (r *UserRepository) HardDeleteUser(ctx context.Context, id int64) error {
	return r.q.HardDeleteUser(ctx, id)
}

func toDomainUser(u gen.User) domain.User {
	user := domain.User{
		ID:        int(u.ID),
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	if u.DeletedAt.Valid {
		user.DeletedAt = &u.DeletedAt.Time
	}
	return user
}
