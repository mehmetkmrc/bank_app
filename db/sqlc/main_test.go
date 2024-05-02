package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const(
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:ITM-2020@localhost:5432/new_banka?sslmode=disable" 
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err :=sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}