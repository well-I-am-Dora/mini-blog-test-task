package comment

import (
	"context"
	mock_comment_service "github.com/martyushova/posts/internal/service/comment/mocks"
	"go.uber.org/mock/gomock"
	"testing"
)

type testDeps struct {
	ctx  context.Context
	ctrl *gomock.Controller

	service *Service

	commentRepository *mock_comment_service.MockcommentRepository
}

func setupTestDeps(t *testing.T) *testDeps {
	deps := &testDeps{
		ctx:  context.Background(),
		ctrl: gomock.NewController(t),
	}
	deps.commentRepository = mock_comment_service.NewMockcommentRepository(deps.ctrl)

	deps.service = NewService(deps.commentRepository)

	return deps
}
