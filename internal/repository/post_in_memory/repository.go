package post_in_memory

import (
	"github.com/hashicorp/go-memdb"
)

// Repository репозиторий постов
type Repository struct {
	db *memdb.MemDB
}

// NewRepository создает экземпляр репозитория постов
func NewRepository(
	db *memdb.MemDB,
) *Repository {
	return &Repository{
		db: db,
	}
}
