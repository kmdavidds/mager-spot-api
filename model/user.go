package model

import "github.com/google/uuid"

type UserRegister struct {
	ID          uuid.UUID `json:"-"`
	Username    string    `json:"username" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
	Password    string    `json:"password" binding:"required,min=8"`
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
