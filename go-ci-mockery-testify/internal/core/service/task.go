package service

import (
	"context"

	"github.com/yehezkiel1086/go-rest-mockery-testify/internal/core/domain"
	"github.com/yehezkiel1086/go-rest-mockery-testify/internal/core/port"
)

type TaskService struct {
	repo port.TaskRepository
}

func NewTaskService(repo port.TaskRepository) *TaskService {
	return &TaskService{
		repo,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	return s.repo.CreateTask(ctx, task)
}

func (s *TaskService) GetTaskByID(ctx context.Context, id uint) (*domain.Task, error) {
	return s.repo.GetTaskByID(ctx, id)
}

func (s *TaskService) GetTasks(ctx context.Context) ([]domain.Task, error) {
	return s.repo.GetTasks(ctx)
}

func (s *TaskService) UpdateTask(ctx context.Context, id uint, task *domain.Task) (*domain.Task, error) {
	// get task
	foundTask, err := s.repo.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// update task
	if task.Description != "" {
		foundTask.Description = task.Description
	}
	if task.Name != "" {
		foundTask.Name = task.Name
	}
	if task.Status != "" {
		foundTask.Status = task.Status
	}

	return s.repo.UpdateTask(ctx, id, task)
}

func (s *TaskService) DeleteTask(ctx context.Context, id uint) error {
	return s.repo.DeleteTask(ctx, id)
}
