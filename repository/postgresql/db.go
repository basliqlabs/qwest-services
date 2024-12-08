package postgresql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Config struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Host     string `koanf:"host"`
	Port     uint   `koanf:"port"`
	DBName   string `koanf:"dbname"`
}

type DB struct {
	db     *sql.DB
	config Config
}

func (pg *DB) Conn() *sql.DB {
	return pg.db
}

func New(config Config) *DB {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.DBName,
	))

	if err != nil {
		panic(fmt.Errorf("failed to open postgresql database: %v", err))
	}

	// TODO: add to config
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &DB{db: db, config: config}
}
