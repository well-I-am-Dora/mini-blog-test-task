package comment_pg

import "github.com/jmoiron/sqlx"

// Repository репозиторий комментариев
type Repository struct {
	db *sqlx.DB
}

// NewRepository создает экземпляр репозитория комментариев
func NewRepository(
	db *sqlx.DB,
) *Repository {
	return &Repository{
		db: db,
	}
}
