package postgresqlrefreshtoken

import "github.com/basliqlabs/qwest-services/internal/repository/postgresql"

type Repository struct {
	db *postgresql.DB
}

func New(db *postgresql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
