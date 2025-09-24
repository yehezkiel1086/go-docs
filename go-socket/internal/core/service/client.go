package service

import (
	"context"
	"go-socket/internal/core/port"
)

type ClientService struct {
	repo port.ClientRepository
}

func InitClientService(repo port.ClientRepository) *ClientService {
	return &ClientService{
		repo: repo,
	}
}

func (cs *ClientService) WriteMessage(ctx context.Context) {
	
}
