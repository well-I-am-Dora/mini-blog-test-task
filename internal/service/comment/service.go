package comment

import "sync"

// Service - сервис комментариев
type Service struct {
	commentRepository commentRepository

	commentNotifier commentNotifier
}

// NewService создает экземпляр сервиса комментариев
func NewService(
	commentRepository commentRepository,
) *Service {
	return &Service{
		commentRepository: commentRepository,
		commentNotifier: commentNotifier{
			mu:            sync.RWMutex{},
			subscriptions: map[string][]*subscription{},
		},
	}
}
