package post_pg

import "github.com/jmoiron/sqlx"

// Repository репозиторий постов
type Repository struct {
	db *sqlx.DB
}

// NewRepository создает экземпляр репозитория постов
func NewRepository(
	db *sqlx.DB,
) *Repository {
	return &Repository{
		db: db,
	}
}
