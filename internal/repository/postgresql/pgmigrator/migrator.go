package pgmigrator

import (
	"database/sql"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"

	"github.com/basliqlabs/qwest-services/internal/repository/postgresql"
)

type Migrator struct {
	dbConfig   postgresql.Config
	migrations *migrate.FileMigrationSource
	dialect    string
}

func NewMigrator(dbConfig postgresql.Config) Migrator {
	migrations := &migrate.FileMigrationSource{
		Dir: "repository/postgresql/migrations",
	}
	return Migrator{
		migrations: migrations,
		dbConfig:   dbConfig,
		dialect:    "postgres",
	}
}

func (m Migrator) Up() {
	db, err := sql.Open(m.dialect, fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		m.dbConfig.Host,
		m.dbConfig.Port,
		m.dbConfig.Username,
		m.dbConfig.Password,
		m.dbConfig.DBName,
	))

	if err != nil {
		panic(fmt.Errorf("failed to open postgresql database: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Up)

	if err != nil {
		panic(fmt.Errorf("can't apply migrations: %v", err))
	}

	fmt.Printf("Applied %d migrations", n)
}

func (m Migrator) Down() {
	db, err := sql.Open(m.dialect, fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		m.dbConfig.Host,
		m.dbConfig.Port,
		m.dbConfig.Username,
		m.dbConfig.Password,
		m.dbConfig.DBName,
	))

	if err != nil {
		panic(fmt.Errorf("failed to open postgresql database: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Down)

	if err != nil {
		panic(fmt.Errorf("can't rollback migrations: %v", err))
	}

	fmt.Printf("Rollbacked %d migrations", n)
}
