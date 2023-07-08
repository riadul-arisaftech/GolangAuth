package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/riad/simple_auth/src/db/sqlc"
	HTTP "github.com/riad/simple_auth/src/http"
	"github.com/riad/simple_auth/src/util"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		fmt.Printf("cannot load config")
	}

	// connPool, err := sql.Open(config.Database.Driver, config.Database.GetDBSource())
	connPool, err := pgxpool.New(context.Background(), config.Database.GetDBSource())
	if err != nil {
		fmt.Printf("cannot connect to db")
	}
	defer connPool.Close()

	runDBMigration(config.Database.MigrateUrl, config.Database.GetDBSource())

	store := db.NewStore(connPool)

	runGinServer(config, store)
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		fmt.Printf("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Printf("failed to run migrate up")
	}

	fmt.Printf("db migrated successfully")
}

func runGinServer(config util.Config, store db.Store) {
	server, err := HTTP.NewServer(config, store)
	if err != nil {
		fmt.Printf("cannot create server")
	}

	err = server.Start(config.Server.HTTPAddress)
	if err != nil {
		fmt.Printf("cannot start server")
	}
}
