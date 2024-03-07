package model

import (
	"time"

	"github.com/google/uuid"
)

type CommentCreate struct {
	ID        uuid.UUID `json:"id"`
	PostID    uuid.UUID `json:"postID"`
	UserID    uuid.UUID `json:"userID"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"-"`
}

type CommentParam struct {
	ID     uuid.UUID `json:"-"`
	PostID uuid.UUID `json:"-"`
}
