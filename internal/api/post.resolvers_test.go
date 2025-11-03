package api

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/martyushova/posts/internal/model"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestPostResolver_Comments(t *testing.T) {
	t.Parallel()

	postID := uuid.NewString()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		deps := setupTestDeps(t)

		deps.commentService.EXPECT().ListComments(gomock.Any(), postID, int32(10), int32(0)).Return(
			[]*model.Comment{
				{
					ID:     uuid.NewString(),
					PostID: postID,
				},
			},
			nil,
		)

		res, err := deps.resolver.Post().Comments(deps.ctx, &model.Post{ID: postID}, 10, 0)
		require.NoError(t, err)
		require.Len(t, res, 1)
	})
	t.Run("error", func(t *testing.T) {
		t.Parallel()

		deps := setupTestDeps(t)

		deps.commentService.EXPECT().ListComments(gomock.Any(), postID, int32(10), int32(0)).Return(nil, fmt.Errorf("error"))

		res, err := deps.resolver.Post().Comments(deps.ctx, &model.Post{ID: postID}, 10, 0)
		require.Error(t, err)
		require.Nil(t, res)
	})
}
