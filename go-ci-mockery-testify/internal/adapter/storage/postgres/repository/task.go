package repository

import (
	"context"

	"github.com/yehezkiel1086/go-rest-mockery-testify/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-rest-mockery-testify/internal/core/domain"
)

type TaskRepository struct {
	db *postgres.DB
}

func NewTaskRepository(db *postgres.DB) (*TaskRepository) {
	return &TaskRepository{
		db,
	}
}

func (r *TaskRepository) CreateTask(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Create(task).Error; err != nil {
		return nil, domain.ErrInternal
	}

	return task, nil
}

func (r *TaskRepository) GetTaskByID(ctx context.Context, id uint) (*domain.Task, error) {
	db := r.db.GetDB()

	var task domain.Task
	if err := db.WithContext(ctx).Where("id = ?", id).First(&task).Error; err != nil {
		return nil, domain.ErrInternal
	}

	return &task, nil
}

func (r *TaskRepository) GetTasks(ctx context.Context) ([]domain.Task, error) {
	db := r.db.GetDB()

	var tasks []domain.Task
	if err := db.WithContext(ctx).Find(&tasks).Error; err != nil {
		return nil, domain.ErrInternal
	}

	return tasks, nil
}

func (r *TaskRepository) UpdateTask(ctx context.Context, id uint, task *domain.Task) (*domain.Task, error) {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Model(&domain.Task{}).Where("id = ?", id).Updates(task).Error; err != nil {
		return nil, domain.ErrInternal
	}

	return task, nil
}

func (r *TaskRepository) DeleteTask(ctx context.Context, id uint) error {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Delete(&domain.Task{}, id).Error; err != nil {
		return domain.ErrInternal
	}

	return nil
}
