package post

import (
	"context"
	"fmt"
	"github.com/martyushova/posts/internal/model"
)

// CreatePost создает новый пост
func (s *Service) CreatePost(ctx context.Context, userID string, post *model.Post) (*model.Post, error) {
	post, err := s.postRepository.CreatePost(ctx, userID, post)
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	return post, nil
}
