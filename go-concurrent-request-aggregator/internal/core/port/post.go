package port

import (
	"github.com/yehezkiel1086/go-concurrent-request-aggregator/internal/core/domain"
)

type PostRepository interface {
	GetPostsByUserId(id string) ([]*domain.Post, error)
}
