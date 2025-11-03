package post

import (
	"errors"
	"github.com/google/uuid"
	"github.com/martyushova/posts/internal/model"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestService_GetPost(t *testing.T) {
	t.Parallel()

	postID := uuid.NewString()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		deps := setupTestDeps(t)

		deps.postRepository.EXPECT().GetPost(gomock.Any(), postID).Return(&model.Post{ID: postID}, nil)

		res, err := deps.service.GetPost(deps.ctx, postID)
		require.NoError(t, err)
		require.NotNil(t, res)
		require.Equal(t, postID, res.ID)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		deps := setupTestDeps(t)

		deps.postRepository.EXPECT().GetPost(gomock.Any(), postID).Return(nil, errors.New("error"))

		res, err := deps.service.GetPost(deps.ctx, postID)
		require.Error(t, err)
		require.Nil(t, res)
	})
}
