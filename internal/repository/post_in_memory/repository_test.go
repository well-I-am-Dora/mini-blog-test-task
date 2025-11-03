package post_in_memory

import (
	"context"
	"github.com/martyushova/posts/internal/utils/test"
	"testing"
)

type testDeps struct {
	ctx context.Context

	repository *Repository
}

func setupTestDeps(t *testing.T) *testDeps {
	deps := &testDeps{
		ctx: context.Background(),
	}

	deps.repository = NewRepository(test.NewMemDB(t))

	return deps
}
