package model

import "time"

// Comment - модель комментария
type Comment struct {
	ID              string
	PostID          string
	UserID          string
	Text            string
	ParentCommentID *string
	Path            string

	CreatedAt time.Time
	UpdatedAt time.Time
}
