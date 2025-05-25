// integration_test.go
package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func TestIntegrationJokeHandler_InsertsIntoDB(t *testing.T) {
	// 1. Ensure env vars match your Docker Compose
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "mysecretpassword")
	os.Setenv("DB_NAME", "jokes_db")

	// 2. Open a fresh DB connection
	dsn := "host=localhost port=5432 user=postgres password=mysecretpassword dbname=jokes_db sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("db open: %v", err)
	}
	defer db.Close()

	// 3. Count rows before
	var before int
	if err := db.QueryRow("SELECT count(*) FROM jokes").Scan(&before); err != nil {
		t.Fatalf("count before: %v", err)
	}

	// 4. Call the handler via HTTP
	srv := setupServer()
	defer srv.Close()

	res, err := http.Get(srv.URL + "/joke")
	if err != nil {
		t.Fatalf("GET /joke error: %v", err)
	}
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", res.StatusCode)
	}

	// 5. Decode and sanity-check the payload
	var j Joke
	if err := json.NewDecoder(res.Body).Decode(&j); err != nil {
		// Note: in your real code you might buffer the body before closing.
	}

	// 6. Count rows after
	var after int
	if err := db.QueryRow("SELECT count(*) FROM jokes").Scan(&after); err != nil {
		t.Fatalf("count after: %v", err)
	}
	if after != before+1 {
		t.Errorf("expected row count to increase by 1 (before=%d, after=%d)", before, after)
	}
}
