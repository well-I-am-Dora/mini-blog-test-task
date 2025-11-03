package comment

import (
	"context"
	"fmt"
	"github.com/martyushova/posts/internal/model"
)

// ListComments получение списка комментариев для конкретного поста по его id
func (s *Service) ListComments(ctx context.Context, postID string, limit int32, offset int32) ([]*model.Comment, error) {
	comments, err := s.commentRepository.ListComments(ctx, postID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("listing comments: %w", err)
	}

	return comments, nil
}
