package post

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/martyushova/posts/internal/model"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestService_CreatePost(t *testing.T) {
	t.Parallel()

	userID := uuid.NewString()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		deps := setupTestDeps(t)

		post := &model.Post{
			ID:     uuid.NewString(),
			UserID: userID,
			Text:   "test",
		}

		deps.postRepository.EXPECT().CreatePost(gomock.Any(), userID, &model.Post{Text: "test"}).Return(post, nil)

		res, err := deps.service.CreatePost(deps.ctx, userID, &model.Post{Text: "test"})
		require.NoError(t, err)
		require.Equal(t, post, res)
	})
	t.Run("error", func(t *testing.T) {
		t.Parallel()

		deps := setupTestDeps(t)

		deps.postRepository.EXPECT().CreatePost(gomock.Any(), userID, &model.Post{Text: "test"}).Return(nil, fmt.Errorf("error"))

		res, err := deps.service.CreatePost(deps.ctx, userID, &model.Post{Text: "test"})
		require.Error(t, err)
		require.Nil(t, res)
	})
}
