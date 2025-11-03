package comment_pg

import (
	"github.com/martyushova/posts/internal/model"
	"github.com/samber/lo"
	"time"
)

// Comment - модель комментария
type Comment struct {
	ID              string  `db:"id"`
	Text            string  `db:"text"`
	UserID          string  `db:"user_id"`
	PostID          string  `db:"post_id"`
	ParentCommentID *string `db:"parent_comment_id"`
	Path            *string `db:"path"`

	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

func toModel(in *Comment) *model.Comment {
	return &model.Comment{
		ID:              in.ID,
		Text:            in.Text,
		UserID:          in.UserID,
		PostID:          in.PostID,
		ParentCommentID: in.ParentCommentID,

		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}

func toModels(in []*Comment) []*model.Comment {
	return lo.Map(
		in,
		func(item *Comment, index int) *model.Comment {
			return toModel(item)
		},
	)
}
