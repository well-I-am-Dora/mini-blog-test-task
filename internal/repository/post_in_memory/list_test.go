package post_in_memory

import (
	"github.com/google/uuid"
	"github.com/martyushova/posts/internal/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRepository_ListPosts(t *testing.T) {
	t.Parallel()

	userID := uuid.NewString()
	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		deps := setupTestDeps(t)
		post := &model.Post{
			ID:     uuid.NewString(),
			UserID: userID,
			Text:   "test",
		}
		_, err := deps.repository.CreatePost(deps.ctx, userID, post)
		require.NoError(t, err)

		res, err := deps.repository.ListPosts(deps.ctx)
		require.NoError(t, err)
		require.Len(t, res, 1)
	})

}
