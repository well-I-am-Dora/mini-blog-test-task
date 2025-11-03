package post_pg

import (
	"context"
	"fmt"
	"github.com/martyushova/posts/internal/model"
)

// ListPosts получить список постов
func (r *Repository) ListPosts(ctx context.Context) ([]*model.Post, error) {
	sql := `SELECT * FROM posts;`
	var posts []*Post
	err := r.db.SelectContext(ctx, &posts, sql)
	if err != nil {
		return nil, fmt.Errorf("listing posts: %w", err)
	}

	return toModels(posts), nil
}
