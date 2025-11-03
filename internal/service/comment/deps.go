package comment

import (
	"context"
	"github.com/martyushova/posts/internal/model"
)

//go:generate mockgen -source $GOFILE -destination mocks/deps.mock.go -package=mock_comment_service

type commentRepository interface {
	CreateComment(ctx context.Context, userID string, comment *model.Comment) (*model.Comment, error)
	ListComments(ctx context.Context, postID string, limit int32, offset int32) ([]*model.Comment, error)
}
