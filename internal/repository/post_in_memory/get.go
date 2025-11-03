package post_in_memory

import (
	"context"
	"errors"
	"fmt"
	"github.com/martyushova/posts/internal/model"
)

// GetPost получить пост по id
func (r *Repository) GetPost(_ context.Context, id string) (*model.Post, error) {
	txn := r.db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("posts", "id", id)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}
	if raw == nil {
		return nil, model.ErrNotFound
	}

	post, ok := raw.(*model.Post)
	if !ok {
		return nil, errors.New("failed to cast post")
	}

	return post, nil
}
