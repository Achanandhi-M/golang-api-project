package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestAPI_MemeHandler(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/meme")
	if err != nil {
		t.Fatalf("GET /meme error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if ct := resp.Header.Get("Content-Type"); !strings.HasPrefix(ct, "text/html") {
		t.Errorf("Expected Content-Type text/html, got %s", ct)
	}

	body, _ := io.ReadAll(resp.Body)
	if !strings.Contains(string(body), "<img") {
		t.Error("Expected <img> tag in meme HTML")
	}
}

func TestAPI_JokeHandler(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/joke")
	if err != nil {
		t.Fatalf("GET /joke error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", resp.StatusCode)
	}

	var joke Joke
	err = json.NewDecoder(resp.Body).Decode(&joke)
	if err != nil {
		t.Errorf("Failed to decode joke response: %v", err)
	}
	if joke.Joke == "" {
		t.Error("Joke text should not be empty")
	}
}
