package post

// Service - сервис постов
type Service struct {
	postRepository postRepository
}

// NewService создает экземпляр сервиса постов
func NewService(
	postRepository postRepository,
) *Service {
	return &Service{
		postRepository: postRepository,
	}
}
