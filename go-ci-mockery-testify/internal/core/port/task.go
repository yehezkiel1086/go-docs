package port

import (
	"context"

	"github.com/yehezkiel1086/go-rest-mockery-testify/internal/core/domain"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task *domain.Task) (*domain.Task, error)
	GetTaskByID(ctx context.Context, id uint) (*domain.Task, error)
	GetTasks(ctx context.Context) ([]domain.Task, error)
	UpdateTask(ctx context.Context, id uint, task *domain.Task) (*domain.Task, error)
	DeleteTask(ctx context.Context, id uint) error
}

type TaskService interface {
	CreateTask(ctx context.Context, task *domain.Task) (*domain.Task, error)
	GetTaskByID(ctx context.Context, id uint) (*domain.Task, error)
	GetTasks(ctx context.Context) ([]domain.Task, error)
	UpdateTask(ctx context.Context, id uint, task *domain.Task) (*domain.Task, error)
	DeleteTask(ctx context.Context, id uint) error
}
