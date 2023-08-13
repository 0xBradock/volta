package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/volta?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		panic("failed")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
