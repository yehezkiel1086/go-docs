package repository

import (
	"context"
	"errors"
	storage "go-socket/internal/adapter/storage/postgres"
	"go-socket/internal/core/domain"
)

type HubRepository struct {
	// db *storage.DB
	Hub *domain.Hub
}

func InitHubRepository(db *storage.DB) *HubRepository {
	return &HubRepository{
		Hub: &domain.Hub{
			Rooms: make(map[string]*domain.Room),
			Register: make(chan *domain.Client),
			Unregister: make(chan *domain.Client),
			Broadcast: make(chan *domain.Message, 5),
		},
	}
}

func (hr *HubRepository) CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error) {
	if _, ok := hr.Hub.Rooms[room.ID]; ok {
		return &domain.Room{}, errors.New("Room already exists.")
	}

	hr.Hub.Rooms[room.ID] = room
	return room, nil
}
