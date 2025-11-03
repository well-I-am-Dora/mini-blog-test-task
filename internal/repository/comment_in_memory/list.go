package comment_in_memory

import (
	"context"
	"fmt"
	"github.com/martyushova/posts/internal/model"
	"sort"
)

// ListComments получить список комментариев
func (r *Repository) ListComments(_ context.Context, postID string, limit int32, offset int32) ([]*model.Comment, error) {
	txn := r.db.Txn(false)
	defer txn.Abort()

	it, err := txn.Get("comments", "post_id", postID)
	if err != nil {
		return nil, fmt.Errorf("failed to list comments: %w", err)
	}

	comments := make([]*model.Comment, 0)
	for obj := it.Next(); obj != nil; obj = it.Next() {
		comments = append(comments, obj.(*model.Comment))
	}

	sort.Slice(comments, func(i, j int) bool {
		return comments[i].Path < comments[j].Path
	})

	return comments[offset:min(int(offset+limit), len(comments))], nil
}
