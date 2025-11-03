package post

import (
	"context"
	"github.com/martyushova/posts/internal/model"
)

// GetPost возвращает пост по id
func (s *Service) GetPost(ctx context.Context, id string) (*model.Post, error) {
	post, err := s.postRepository.GetPost(ctx, id)
	if err != nil {
		return nil, err
	}

	return post, nil
}
