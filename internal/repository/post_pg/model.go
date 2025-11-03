package post_pg

import (
	"github.com/martyushova/posts/internal/model"
	"github.com/samber/lo"
	"time"
)

// Post - модель поста
type Post struct {
	ID              string `db:"id"`
	Text            string `db:"text"`
	UserID          string `db:"user_id"`
	CommentsAllowed bool   `db:"comments_allowed"`

	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

func toModel(in *Post) *model.Post {
	return &model.Post{
		ID:              in.ID,
		Text:            in.Text,
		UserID:          in.UserID,
		CommentsAllowed: in.CommentsAllowed,

		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}

func toModels(in []*Post) []*model.Post {
	return lo.Map(
		in,
		func(item *Post, index int) *model.Post {
			return toModel(item)
		},
	)
}
