package postgresqluser

import "github.com/basliqlabs/qwest-services/internal/repository/postgresql"

type DB struct {
	db *postgresql.DB
}

func New(db *postgresql.DB) *DB {
	return &DB{
		db: db,
	}
}
