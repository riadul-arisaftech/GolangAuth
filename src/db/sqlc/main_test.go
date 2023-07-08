package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"

	"github.com/riad/simple_auth/src/util"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../../.")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	// testDB, err := sql.Open(config.Database.Driver, config.Database.GetDBSource())
	connPool, err := pgxpool.New(context.Background(), config.Database.GetDBSource())
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testStore = NewStore(connPool)

	os.Exit(m.Run())
}
