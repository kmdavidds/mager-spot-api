package model

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type UserRegister struct {
	ID       uuid.UUID `json:"-"`
	Username string    `json:"username" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required,min=8"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserParam struct {
	ID       uuid.UUID `json:"-"`
	Username string    `json:"-"`
}

type UserUpdates struct {
	DisplayName string `json:"displayName"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

type PhotoUpdate struct {
	UserID    uuid.UUID             `json:"-"`
	PhotoLink string                `json:"-"`
	Image     *multipart.FileHeader `form:"image" binding:"required"`
}
