package comment

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/martyushova/posts/internal/model"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestService_CreateComment(t *testing.T) {
	t.Parallel()

	userID := uuid.NewString()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		deps := setupTestDeps(t)

		comment := &model.Comment{
			ID:     uuid.NewString(),
			UserID: userID,
			Text:   "test comment",
		}

		deps.commentRepository.EXPECT().CreateComment(gomock.Any(), userID, &model.Comment{Text: "test comment"}).Return(comment, nil)

		res, err := deps.service.CreateComment(deps.ctx, userID, &model.Comment{Text: "test comment"})
		require.NoError(t, err)
		require.Equal(t, comment, res)
	})
	t.Run("error", func(t *testing.T) {
		t.Parallel()

		deps := setupTestDeps(t)

		deps.commentRepository.EXPECT().CreateComment(gomock.Any(), userID, &model.Comment{Text: "test comment"}).Return(nil, fmt.Errorf("error"))

		res, err := deps.service.CreateComment(deps.ctx, userID, &model.Comment{Text: "test comment"})
		require.Error(t, err)
		require.Nil(t, res)
	})
}
