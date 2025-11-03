package post

import (
	"context"
	"github.com/martyushova/posts/internal/model"
)

//go:generate mockgen -source $GOFILE -destination mocks/deps.mock.go -package=mock_post_service

type postRepository interface {
	CreatePost(ctx context.Context, userID string, post *model.Post) (*model.Post, error)
	ListPosts(ctx context.Context) ([]*model.Post, error)
	GetPost(ctx context.Context, postID string) (*model.Post, error)
}
