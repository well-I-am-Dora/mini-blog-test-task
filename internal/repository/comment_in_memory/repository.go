package comment_in_memory

import (
	"github.com/hashicorp/go-memdb"
)

// Repository репозиторий комментариев
type Repository struct {
	db *memdb.MemDB
}

// NewRepository создает экземпляр репозитория комментариев
func NewRepository(
	db *memdb.MemDB,
) *Repository {
	return &Repository{
		db: db,
	}
}
