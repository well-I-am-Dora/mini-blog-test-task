package comment

import (
	"context"
	"github.com/martyushova/posts/internal/model"
	"sync"
)

type subscription struct {
	id  string
	ctx context.Context
	ch  chan *model.Comment
}

type commentNotifier struct {
	mu sync.RWMutex

	// Ключ - postID
	// Значение - список подписок
	subscriptions map[string][]*subscription
}

// unsubscribe удаляет подписку
func (c *commentNotifier) unsubscribe(postID, subscriptionID string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	subs, exists := c.subscriptions[postID]
	if !exists {
		return false
	}

	for i, sub := range subs {
		if sub.id == subscriptionID {
			c.subscriptions[postID] = append(subs[:i], subs[i+1:]...)

			if len(c.subscriptions[postID]) == 0 {
				delete(c.subscriptions, postID)
			}
			return true
		}
	}
	return false
}

// SubscribeOnPost подписаться на комментарии поста
func (s *Service) SubscribeOnPost(ctx context.Context, postID string) (<-chan *model.Comment, error) {
	ch := make(chan *model.Comment)
	sub := &subscription{
		ctx: ctx,
		ch:  ch,
	}

	s.commentNotifier.mu.Lock()
	s.commentNotifier.subscriptions[postID] = append(s.commentNotifier.subscriptions[postID], sub)
	s.commentNotifier.mu.Unlock()

	return ch, nil
}
