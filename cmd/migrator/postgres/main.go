package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/basliqlabs/qwest-services/config"
	"github.com/basliqlabs/qwest-services/repository/postgresql/pgmigrator"
)

func main() {
	cfg := config.Load("config.yml")

	// TODO: add limit flag
	up := flag.Bool("up", false, "Migrate the database up")
	down := flag.Bool("down", false, "Migrate the database down")
	flag.Parse()

	migrator := pgmigrator.NewMigrator(cfg.Repository.Postgres)

	if *up {
		migrator.Up()
	} else if *down {
		migrator.Down()
	} else {
		fmt.Println("Please specify a migration direction with -up or -down")
		flag.Usage()
		os.Exit(1)
	}
}
