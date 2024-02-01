package repository

import "github.com/jmoiron/sqlx"

type Create interface {
	CreateUser() (int, float64, error)
}

type Repository struct {
	Create
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Create: NewCreatePostgres(db),
	}
}
