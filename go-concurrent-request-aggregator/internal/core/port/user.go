package port

import (
	"github.com/yehezkiel1086/go-concurrent-request-aggregator/internal/core/domain"
)

type UserRepository interface {
	GetUser(id string) (*domain.User, error)
}
