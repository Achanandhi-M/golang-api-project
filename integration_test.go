package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func TestIntegration_JokeHandler_InsertsIntoDB(t *testing.T) {
	// Load environment variables from .env file
	_ = godotenv.Load(".env")

	// Connect to Postgres using env vars
	dsn := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	dsn = formatDSNFromEnv(dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("DB open error: %v", err)
	}
	defer db.Close()

	// Count jokes before request
	var before int
	if err := db.QueryRow("SELECT COUNT(*) FROM jokes").Scan(&before); err != nil {
		t.Fatalf("Count before insert: %v", err)
	}

	// Call /joke API
	srv := setupServer()
	defer srv.Close()

	res, err := http.Get(srv.URL + "/joke")
	if err != nil {
		t.Fatalf("GET /joke error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", res.StatusCode)
	}

	// Count jokes after insert
	var after int
	if err := db.QueryRow("SELECT COUNT(*) FROM jokes").Scan(&after); err != nil {
		t.Fatalf("Count after insert: %v", err)
	}
	if after != before+1 {
		t.Errorf("Expected row count to increase by 1 (before=%d, after=%d)", before, after)
	}
}

// Helper to format DSN using env vars
func formatDSNFromEnv(template string) string {
	return fmt.Sprintf(template,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
