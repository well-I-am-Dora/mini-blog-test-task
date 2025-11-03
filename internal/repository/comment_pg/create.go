package comment_pg

import (
	"context"
	"fmt"
	"github.com/martyushova/posts/internal/model"
)

// CreateComment создает новый комментарий
func (r *Repository) CreateComment(ctx context.Context, userID string, comment *model.Comment) (*model.Comment, error) {
	var pgComment Comment

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	insertSQL := `INSERT INTO comments (user_id, text, post_id, parent_comment_id) VALUES ($1, $2, $3, $4) RETURNING *;`
	err = tx.QueryRowxContext(ctx, insertSQL, userID, comment.Text, comment.PostID, comment.ParentCommentID).StructScan(&pgComment)
	if err != nil {
		return nil, fmt.Errorf("insert comment: %w", err)
	}

	var path string
	// Если есть родитель, получаем его путь
	if pgComment.ParentCommentID != nil {
		parentComment := &Comment{}
		selectSQL := `SELECT path FROM comments WHERE id = $1;`
		err = tx.GetContext(ctx, parentComment, selectSQL, *pgComment.ParentCommentID)
		if err != nil {
			return nil, fmt.Errorf("get parent comment: %w", err)
		}

		if parentComment.Path == nil {
			return nil, fmt.Errorf("parent comment path is nil")
		}
		path = *parentComment.Path + "." + pgComment.ID
	} else {
		path = pgComment.ID
	}

	// Обновляем путь
	updateSQL := `UPDATE comments SET path = $1 WHERE id = $2 RETURNING *;`
	err = tx.QueryRowxContext(ctx, updateSQL, path, pgComment.ID).StructScan(&pgComment)
	if err != nil {
		return nil, fmt.Errorf("update comment path: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("commit transaction: %w", err)
	}

	return toModel(&pgComment), nil
}
