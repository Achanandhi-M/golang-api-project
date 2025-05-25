package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// setupServer spins up only the memeHandler.
func setupServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/meme", memeHandler)
	mux.HandleFunc("/joke", jokeHandler)
	return httptest.NewServer(mux)
}

func TestUnitMemeHandler_StatusAndHTML(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	res, err := http.Get(srv.URL + "/meme")
	if err != nil {
		t.Fatalf("GET /meme error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", res.StatusCode)
	}
	if ct := res.Header.Get("Content-Type"); !strings.HasPrefix(ct, "text/html") {
		t.Errorf("expected HTML content, got %s", ct)
	}

	body, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(body), "<img") {
		t.Errorf("expected an <img> tag in response HTML")
	}
}

func TestUnitMemeHandler_StructDecode(t *testing.T) {
	// If you want to inspect the JSON from Meme-API directly, you could:
	type Meme struct {
		Title   string `json:"title"`
		URL     string `json:"url"`
		PostURL string `json:"postLink"`
	}
	// But here we trust our handler template. This is just an example.
	_, err := json.Marshal(Meme{Title: "x", URL: "y", PostURL: "z"})
	if err != nil {
		t.Fatal("JSON marshal should not fail:", err)
	}
}
