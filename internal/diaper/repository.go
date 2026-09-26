package diaper

import "github.com/dakotaodev/cradle/internal/db"

type Repository struct {
	db db.Querier
}

func NewRepository(database db.Querier) *Repository {
	return &Repository{db: database}
}
