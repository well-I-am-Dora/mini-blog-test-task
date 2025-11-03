package post

import (
	"errors"
	"github.com/martyushova/posts/internal/model"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestService_ListPosts(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		deps := setupTestDeps(t)

		deps.postRepository.EXPECT().ListPosts(gomock.Any()).Return([]*model.Post{{ID: "123"}}, nil)

		res, err := deps.service.ListPosts(deps.ctx)
		require.NoError(t, err)
		require.Len(t, res, 1)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		deps := setupTestDeps(t)

		deps.postRepository.EXPECT().ListPosts(gomock.Any()).Return(nil, errors.New("error"))

		res, err := deps.service.ListPosts(deps.ctx)
		require.Error(t, err)
		require.Nil(t, res)
	})
}
