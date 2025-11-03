package api

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/martyushova/posts/internal/model"
	graphql "github.com/martyushova/posts/internal/model/graphql"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestMutationResolver_CreatePost(t *testing.T) {
	t.Parallel()

	post := &model.Post{
		ID:   uuid.NewString(),
		Text: "test post",
	}

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		deps := setupTestDeps(t)

		deps.postService.EXPECT().CreatePost(gomock.Any(), gomock.Any(), gomock.Any()).Return(post, nil)

		res, err := deps.resolver.Mutation().CreatePost(deps.ctx, graphql.PostInput{Text: post.Text})
		require.NoError(t, err)
		require.NotNil(t, res)
	})
	t.Run("error", func(t *testing.T) {
		t.Parallel()

		deps := setupTestDeps(t)

		deps.postService.EXPECT().CreatePost(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("error"))

		res, err := deps.resolver.Mutation().CreatePost(deps.ctx, graphql.PostInput{Text: post.Text})
		require.Error(t, err)
		require.Nil(t, res)
	})
}
