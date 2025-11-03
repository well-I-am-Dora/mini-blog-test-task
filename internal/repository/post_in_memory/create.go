package post_in_memory

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/martyushova/posts/internal/model"
	"time"
)

// CreatePost создает новый пост
func (r *Repository) CreatePost(_ context.Context, userID string, post *model.Post) (*model.Post, error) {
	post.ID = uuid.NewString()
	post.UserID = userID
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	txn := r.db.Txn(true)
	err := txn.Insert("posts", post)
	if err != nil {
		return nil, fmt.Errorf("failed to insert post: %w", err)
	}

	txn.Commit()
	return post, nil
}
