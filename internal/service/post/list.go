package post

import (
	"context"
	"github.com/martyushova/posts/internal/model"
)

// ListPosts возвращает список постов
func (s *Service) ListPosts(ctx context.Context) ([]*model.Post, error) {
	posts, err := s.postRepository.ListPosts(ctx)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
