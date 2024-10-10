package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:mysecretpassword@localhost:5433/golangpro?sslmode=disable"
)

var testqueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can not connect to db", err)
	}

	testqueries = New(testDB)
	os.Exit(m.Run())
}
