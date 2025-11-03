package model

import "time"

// Post - модель поста
type Post struct {
	ID              string
	Text            string
	UserID          string
	CommentsAllowed bool

	CreatedAt time.Time
	UpdatedAt time.Time
}
