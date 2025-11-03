package api

import (
	"context"
	"github.com/martyushova/posts/internal/model"
)

//go:generate mockgen -source $GOFILE -destination mocks/deps.mock.go -package=mock_api

type postService interface {
	CreatePost(ctx context.Context, userID string, post *model.Post) (*model.Post, error)
	ListPosts(ctx context.Context) ([]*model.Post, error)
	GetPost(ctx context.Context, postID string) (*model.Post, error)
}

type commentService interface {
	SubscribeOnPost(ctx context.Context, postID string) (<-chan *model.Comment, error)
	CreateComment(ctx context.Context, userID string, comment *model.Comment) (*model.Comment, error)
	ListComments(ctx context.Context, postID string, limit int32, offset int32) ([]*model.Comment, error)
}
