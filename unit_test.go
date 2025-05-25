package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

// setupServer loads DB and mounts routes to simulate API server
func setupServer() *httptest.Server {
	_ = godotenv.Load(".env") // load env automatically for tests
	setupDB()                 // connect and migrate DB
	mux := http.NewServeMux()
	mux.HandleFunc("/meme", memeHandler)
	mux.HandleFunc("/joke", jokeHandler)
	return httptest.NewServer(mux)
}

// TestUnit_MemeHandler_ValidHTML ensures /meme returns a proper HTML with an <img> tag
func TestUnit_MemeHandler_ValidHTML(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	res, err := http.Get(srv.URL + "/meme")
	if err != nil {
		t.Fatalf("GET /meme failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", res.StatusCode)
	}
	if ct := res.Header.Get("Content-Type"); !strings.HasPrefix(ct, "text/html") {
		t.Errorf("Expected HTML content, got %s", ct)
	}

	body, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(body), "<img") {
		t.Errorf("Expected an <img> tag in the response")
	}
}

// TestUnit_Meme_StructJSON verifies that our Meme struct can marshal properly (pure unit test)
func TestUnit_Meme_StructJSON(t *testing.T) {
	type Meme struct {
		Title   string `json:"title"`
		URL     string `json:"url"`
		PostURL string `json:"postLink"`
	}
	_, err := json.Marshal(Meme{Title: "x", URL: "y", PostURL: "z"})
	if err != nil {
		t.Fatal("JSON marshal should not fail:", err)
	}
}
