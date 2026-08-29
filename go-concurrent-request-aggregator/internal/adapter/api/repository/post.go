package repository

import (
	"encoding/json"
	"fmt"

	"github.com/yehezkiel1086/go-concurrent-request-aggregator/internal/adapter/api"
	"github.com/yehezkiel1086/go-concurrent-request-aggregator/internal/core/domain"
)

type PostRepository struct {
	api *api.API
}

func NewPostRepository(api *api.API) *PostRepository {
	return &PostRepository{
		api: api,
	}
}

func (r *PostRepository) GetPostsByUserId(id string) ([]*domain.Post, error) {
	path := fmt.Sprintf("/posts?userId=%s", id)
	resp, err := r.api.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var posts []*domain.Post
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		return nil, err
	}

	return posts, nil
}
