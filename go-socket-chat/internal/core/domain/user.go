package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `json:"username" gorm:"size:255;not null;unique"`
	Email    string `json:"email" gorm:"size:255;email;not null;unique"`
	Password string `json:"password" gorm:"size:255;not null"`
}

type CreateUserReq struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserRes struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginUserReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserRes struct {
	AccessToken string
	Username    string `json:"username"`
}
