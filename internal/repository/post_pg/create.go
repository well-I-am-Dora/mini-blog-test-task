package post_pg

import (
	"context"
	"fmt"
	"github.com/martyushova/posts/internal/model"
)

// CreatePost создает новый пост
func (r *Repository) CreatePost(ctx context.Context, userID string, post *model.Post) (*model.Post, error) {
	sql := `INSERT INTO posts (user_id, text, comments_allowed) VALUES ($1, $2, $3) RETURNING *;`

	pgPost := Post{}
	err := r.db.QueryRowxContext(ctx, sql, userID, post.Text, post.CommentsAllowed).StructScan(&pgPost)
	if err != nil {
		return nil, fmt.Errorf("create post: %w", err)
	}

	return toModel(&pgPost), nil
}
