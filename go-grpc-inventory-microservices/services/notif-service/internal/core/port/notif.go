package port

import (
	"context"

	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/notif-service/internal/core/domain"
)

type NotifRepository interface {
	CreateNotification(ctx context.Context, notif *domain.Notification) (*domain.Notification, error)
	GetNotificationByID(ctx context.Context, id uint) (*domain.Notification, error)
	GetNotificationsByUserID(ctx context.Context, userID uint) ([]domain.Notification, error)
	GetUnreadNotificationsByUserID(ctx context.Context, userID uint) ([]domain.Notification, error)
	MarkAsRead(ctx context.Context, id uint) error
	MarkAllAsRead(ctx context.Context, userID uint) error
	DeleteNotification(ctx context.Context, id uint) error
}

type NotifService interface {
	CreateNotification(ctx context.Context, notif *domain.Notification) (*domain.Notification, error)
	GetNotificationByID(ctx context.Context, id uint) (*domain.Notification, error)
	GetUserNotifications(ctx context.Context, userID uint) ([]domain.Notification, error)
	GetUnreadNotifications(ctx context.Context, userID uint) ([]domain.Notification, error)
	MarkAsRead(ctx context.Context, id uint) error
	MarkAllAsRead(ctx context.Context, userID uint) error
	DeleteNotification(ctx context.Context, id uint) error
}
