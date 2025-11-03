package post_pg

import (
	"context"
	"fmt"
	"github.com/martyushova/posts/internal/model"
)

// GetPost получить пост по id
func (r *Repository) GetPost(ctx context.Context, postID string) (*model.Post, error) {
	req := `SELECT * FROM posts WHERE id = $1;`

	post := &Post{}
	err := r.db.GetContext(ctx, post, req, postID)
	if err != nil {
		return nil, fmt.Errorf("get post: %w", err)
	}

	return toModel(post), nil
}
