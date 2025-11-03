package comment

import (
	"errors"
	"github.com/google/uuid"
	"github.com/martyushova/posts/internal/model"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestService_ListComments(t *testing.T) {
	t.Parallel()

	postID := uuid.NewString()
	limit := int32(5)
	offset := int32(0)
	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		deps := setupTestDeps(t)

		deps.commentRepository.EXPECT().ListComments(gomock.Any(), postID, limit, offset).Return([]*model.Comment{{ID: "123d", UserID: "321s", Text: "vozmite na rabotu"}}, nil)

		res, err := deps.service.ListComments(deps.ctx, postID, limit, offset)
		require.NoError(t, err)
		require.Len(t, res, 1)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		deps := setupTestDeps(t)

		deps.commentRepository.EXPECT().ListComments(gomock.Any(), postID, limit, offset).Return(nil, errors.New("error"))

		res, err := deps.service.ListComments(deps.ctx, postID, limit, offset)
		require.Error(t, err)
		require.Nil(t, res)
	})
}
