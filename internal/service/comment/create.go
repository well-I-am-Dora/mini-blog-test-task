package comment

import (
	"context"
	"fmt"
	"github.com/martyushova/posts/internal/model"
	"sync"
)

// CreateComment создает новый комментарий
func (s *Service) CreateComment(ctx context.Context, userID string, comment *model.Comment) (*model.Comment, error) {
	comment, err := s.commentRepository.CreateComment(ctx, userID, comment)
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	// Асинхронные уведомления
	s.commentNotifier.mu.RLock()
	subs := s.commentNotifier.subscriptions[comment.PostID]
	s.commentNotifier.mu.RUnlock()
	wg := &sync.WaitGroup{}
	for _, sub := range subs {
		wg.Go(func() {
			select {
			case <-sub.ctx.Done():
				s.commentNotifier.unsubscribe(comment.PostID, sub.id)
			default:
				sub.ch <- comment
			}
		})
	}
	wg.Wait()

	return comment, nil
}
