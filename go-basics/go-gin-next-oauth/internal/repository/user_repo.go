package repository

import (
	"go-gin-next-oauth/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) CreateOrUpdate(user *model.User) (*model.User, error) {
	var existing model.User
	result := r.db.Where("email = ?", user.Email).First(&existing)

	if result.Error == gorm.ErrRecordNotFound {
		if err := r.db.Create(user).Error; err != nil {
			return nil, err
		}
		return user, nil
	}

	existing.Name = user.Name
	existing.Picture = user.Picture
	if err := r.db.Save(&existing).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}
