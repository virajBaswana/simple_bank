package db

import (
	"database/sql"
	"log"
	"os"
	db "simple_bank/db/sqlc"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *db.Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:12345@127.0.0.1:9876/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Could not connect to db:", err)
	}
	testQueries = db.New(conn)
	os.Exit(m.Run())
}
