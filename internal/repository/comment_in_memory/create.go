package comment_in_memory

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/martyushova/posts/internal/model"
	"time"
)

// CreateComment создает новый комментарий
func (r *Repository) CreateComment(_ context.Context, userID string, comment *model.Comment) (*model.Comment, error) {
	comment.ID = uuid.NewString()
	comment.UserID = userID
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()
	txn := r.db.Txn(true)
	if comment.ParentCommentID != nil {
		parentComment, err := txn.First("comment", "id", *comment.ParentCommentID)
		if err != nil {
			return nil, fmt.Errorf("find parent comment: %w", err)
		}
		comment.Path = parentComment.(*model.Comment).Path + "." + comment.ID
	} else {
		comment.Path = comment.ID
	}

	err := txn.Insert("comments", comment)
	if err != nil {
		return nil, fmt.Errorf("failed to insert comment: %w", err)
	}

	txn.Commit()
	return comment, nil
}
