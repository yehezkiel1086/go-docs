package repository

import (
	"encoding/json"
	"fmt"

	"github.com/yehezkiel1086/go-concurrent-request-aggregator/internal/adapter/api"
	"github.com/yehezkiel1086/go-concurrent-request-aggregator/internal/core/domain"
)

type UserRepository struct {
	api *api.API
}

func NewUserRepository(api *api.API) *UserRepository {
	return &UserRepository{
		api: api,
	}
}

func (r *UserRepository) GetUser(id string) (*domain.User, error) {
	path := fmt.Sprintf("/users/%s", id)
	resp, err := r.api.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user domain.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
