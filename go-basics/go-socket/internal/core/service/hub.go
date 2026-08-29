package service

import (
	"context"
	"go-socket/internal/core/domain"
	"go-socket/internal/core/port"
)

type HubService struct {
	repo port.HubRepository
}

func InitHubService(repo port.HubRepository) *HubService {
	return &HubService{
		repo: repo,
	}
}

func (hs *HubService) CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error) {
	res, err := hs.repo.CreateRoom(ctx, room)
	if err != nil {
		return &domain.Room{}, err
	}

	return res, nil
}
