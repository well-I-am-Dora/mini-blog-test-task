package api

import (
	"context"
	mock_api "github.com/martyushova/posts/internal/api/mocks"
	"go.uber.org/mock/gomock"
	"testing"
)

type testDeps struct {
	ctx  context.Context
	ctrl *gomock.Controller

	resolver *Resolver

	postService    *mock_api.MockpostService
	commentService *mock_api.MockcommentService
}

func setupTestDeps(t *testing.T) *testDeps {
	deps := &testDeps{
		ctx:  context.Background(),
		ctrl: gomock.NewController(t),
	}
	deps.postService = mock_api.NewMockpostService(deps.ctrl)
	deps.commentService = mock_api.NewMockcommentService(deps.ctrl)

	deps.resolver = NewResolver(deps.postService, deps.commentService)

	return deps
}
