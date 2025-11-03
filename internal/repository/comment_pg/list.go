package comment_pg

import (
	"context"
	"fmt"
	"github.com/martyushova/posts/internal/model"
)

// ListComments получить список комментариев
func (r *Repository) ListComments(ctx context.Context, postID string, limit int32, offset int32) ([]*model.Comment, error) {
	sql := `SELECT * FROM comments WHERE post_id = $1 ORDER BY path LIMIT $2 OFFSET $3;`
	var comments []*Comment
	err := r.db.SelectContext(ctx, &comments, sql, postID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("listing comments: %w", err)
	}

	return toModels(comments), nil
}
