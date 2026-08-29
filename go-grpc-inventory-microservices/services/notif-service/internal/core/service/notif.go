package service

import (
	"context"

	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/notif-service/internal/core/domain"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/notif-service/internal/core/port"
)

type NotifService struct {
	repo port.NotifRepository
}

func NewNotifService(repo port.NotifRepository) *NotifService {
	return &NotifService{
		repo: repo,
	}
}

func (s *NotifService) CreateNotification(ctx context.Context, notif *domain.Notification) (*domain.Notification, error) {
	return s.repo.CreateNotification(ctx, notif)
}

func (s *NotifService) GetNotificationByID(ctx context.Context, id uint) (*domain.Notification, error) {
	return s.repo.GetNotificationByID(ctx, id)
}

func (s *NotifService) GetUserNotifications(ctx context.Context, userID uint) ([]domain.Notification, error) {
	return s.repo.GetNotificationsByUserID(ctx, userID)
}

func (s *NotifService) GetUnreadNotifications(ctx context.Context, userID uint) ([]domain.Notification, error) {
	return s.repo.GetUnreadNotificationsByUserID(ctx, userID)
}

func (s *NotifService) MarkAsRead(ctx context.Context, id uint) error {
	return s.repo.MarkAsRead(ctx, id)
}

func (s *NotifService) MarkAllAsRead(ctx context.Context, userID uint) error {
	return s.repo.MarkAllAsRead(ctx, userID)
}

func (s *NotifService) DeleteNotification(ctx context.Context, id uint) error {
	return s.repo.DeleteNotification(ctx, id)
}
