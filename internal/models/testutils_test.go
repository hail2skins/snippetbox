package models

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func newTestDB(t *testing.T) *sql.DB {
	// Establish a sql.DB connection pool for our test database.
	db, err := sql.Open("postgres", "user=test_web password=pass dbname=test_snippetbox sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	// Read the setup SQL script from the file and execute the statements, closing
	// the connection pool and calling t.Fatal() in the event of an error.
	script, err := os.ReadFile("./testdata/setup.sql")
	if err != nil {
		db.Close()
		t.Fatal(err)
	}
	_, err = db.Exec(string(script))
	if err != nil {
		db.Close()
		t.Fatal(err)
	}

	// Use t.Cleanup() to register a function *which will automatically be
	// called by Go when the current test (or sub-test) which calls newTestDB()
	// has finished*. In this function we read and execute the teardown script,
	// and close the database connection pool.
	t.Cleanup(func() {
		defer db.Close()

		script, err := os.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}
	})

	// Return the database connection pool.
	return db
}
