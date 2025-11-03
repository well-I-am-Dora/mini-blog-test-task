package post_in_memory

import (
	"context"
	"fmt"
	"github.com/martyushova/posts/internal/model"
)

// ListPosts получить список постов
func (r *Repository) ListPosts(_ context.Context) ([]*model.Post, error) {
	txn := r.db.Txn(false)
	defer txn.Abort()

	it, err := txn.Get("posts", "id")
	if err != nil {
		return nil, fmt.Errorf("failed to list posts: %w", err)
	}

	posts := make([]*model.Post, 0)
	for obj := it.Next(); obj != nil; obj = it.Next() {
		posts = append(posts, obj.(*model.Post))
	}

	return posts, nil
}
