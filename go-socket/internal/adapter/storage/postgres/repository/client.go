package repository

import (
	"context"
	"go-socket/internal/core/domain"
)

type ClientRepository struct {
	// db *storage.DB
	Client *domain.Client
	Message *domain.Message
}

func InitClientRepository() *ClientRepository {
	return &ClientRepository{
		Client: &domain.Client{
			Message: make(chan *domain.Message),
		},
		Message: &domain.Message{},
	}
}

func WriteMessage(ctx context.Context) {
	// defer func() {
	// 	c.Conn.Close()
	// }()

	// for {
	// 	message, ok := <-c.Message
	// 	if !ok {
	// 		return
	// 	}

	// 	c.Conn.WriteJSON(message)
	// }
}
