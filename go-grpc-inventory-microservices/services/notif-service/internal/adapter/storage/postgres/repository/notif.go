package repository

import (
	"context"

	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/notif-service/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/notif-service/internal/core/domain"
)

type NotifRepository struct {
	db *postgres.DB
}

func NewNotifRepository(db *postgres.DB) *NotifRepository {
	return &NotifRepository{
		db: db,
	}
}

func (r *NotifRepository) CreateNotification(ctx context.Context, notif *domain.Notification) (*domain.Notification, error) {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Create(notif).Error; err != nil {
		return nil, err
	}

	return notif, nil
}

func (r *NotifRepository) GetNotificationByID(ctx context.Context, id uint) (*domain.Notification, error) {
	db := r.db.GetDB()

	var notif domain.Notification
	if err := db.WithContext(ctx).First(&notif, id).Error; err != nil {
		return nil, err
	}

	return &notif, nil
}

func (r *NotifRepository) GetNotificationsByUserID(ctx context.Context, userID uint) ([]domain.Notification, error) {
	db := r.db.GetDB()

	var notifs []domain.Notification
	if err := db.WithContext(ctx).Where("user_id = ?", userID).Order("created_at DESC").Find(&notifs).Error; err != nil {
		return nil, err
	}

	return notifs, nil
}

func (r *NotifRepository) GetUnreadNotificationsByUserID(ctx context.Context, userID uint) ([]domain.Notification, error) {
	db := r.db.GetDB()

	var notifs []domain.Notification
	if err := db.WithContext(ctx).Where("user_id = ? AND is_read = ?", userID, false).Order("created_at DESC").Find(&notifs).Error; err != nil {
		return nil, err
	}

	return notifs, nil
}

func (r *NotifRepository) MarkAsRead(ctx context.Context, id uint) error {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Model(&domain.Notification{}).Where("id = ?", id).Update("is_read", true).Error; err != nil {
		return err
	}

	return nil
}

func (r *NotifRepository) MarkAllAsRead(ctx context.Context, userID uint) error {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Model(&domain.Notification{}).Where("user_id = ?", userID).Update("is_read", true).Error; err != nil {
		return err
	}

	return nil
}

func (r *NotifRepository) DeleteNotification(ctx context.Context, id uint) error {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Delete(&domain.Notification{}, id).Error; err != nil {
		return err
	}

	return nil
}
