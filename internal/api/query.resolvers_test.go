package api

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/martyushova/posts/internal/model"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestQueryResolver_Posts(t *testing.T) {
	t.Parallel()

	postID := uuid.NewString()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		deps := setupTestDeps(t)

		deps.postService.EXPECT().ListPosts(gomock.Any()).Return([]*model.Post{{ID: postID}}, nil)

		res, err := deps.resolver.Query().Posts(deps.ctx)
		require.NoError(t, err)
		require.Len(t, res, 1)
	})
	t.Run("error", func(t *testing.T) {
		t.Parallel()

		deps := setupTestDeps(t)

		deps.postService.EXPECT().ListPosts(gomock.Any()).Return(nil, fmt.Errorf("error"))

		res, err := deps.resolver.Query().Posts(deps.ctx)
		require.Error(t, err)
		require.Nil(t, res)
	})
}

func TestQueryResolver_Post(t *testing.T) {
	t.Parallel()

	postID := uuid.NewString()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		deps := setupTestDeps(t)

		deps.postService.EXPECT().GetPost(gomock.Any(), postID).Return(&model.Post{ID: postID}, nil)

		res, err := deps.resolver.Query().Post(deps.ctx, postID)
		require.NoError(t, err)
		require.Equal(t, postID, res.ID)
	})
	t.Run("error", func(t *testing.T) {
		t.Parallel()

		deps := setupTestDeps(t)

		deps.postService.EXPECT().GetPost(gomock.Any(), postID).Return(nil, fmt.Errorf("error"))

		res, err := deps.resolver.Query().Post(deps.ctx, postID)
		require.Error(t, err)
		require.Nil(t, res)
	})
}
