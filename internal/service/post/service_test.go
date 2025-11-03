package post

import (
	"context"
	mock_post_service "github.com/martyushova/posts/internal/service/post/mocks"
	"go.uber.org/mock/gomock"
	"testing"
)

type testDeps struct {
	ctx  context.Context
	ctrl *gomock.Controller

	service *Service

	postRepository *mock_post_service.MockpostRepository
}

func setupTestDeps(t *testing.T) *testDeps {
	deps := &testDeps{
		ctx:  context.Background(),
		ctrl: gomock.NewController(t),
	}
	deps.postRepository = mock_post_service.NewMockpostRepository(deps.ctrl)

	deps.service = NewService(deps.postRepository)

	return deps
}
