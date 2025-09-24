package port

import (
	"context"
	"go-socket/internal/core/domain"
)

type HubRepository interface {
	CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error)
}

type HubService interface {
	CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error)
}
