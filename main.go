package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Joke struct {
	Category string `json:"category"`
	Type     string `json:"type"`
	Joke     string `json:"joke"`
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
}

type Meme struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	PostURL string `json:"postLink"`
}

var db *sql.DB

func setupDB() {
	_ = godotenv.Load(".env") // loads from .env file

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("DB open: %v", err)
	}

	for i := 0; i < 10; i++ {
		if err = db.Ping(); err == nil {
			break
		}
		log.Printf("Waiting for DB... (%d/10)", i+1)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("DB never ready: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS jokes (
			id SERIAL PRIMARY KEY,
			category TEXT,
			type TEXT,
			joke TEXT,
			setup TEXT,
			delivery TEXT
		);
	`)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}

func fetchJoke() (*Joke, error) {
	resp, err := http.Get("https://v2.jokeapi.dev/joke/Programming?type=single")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var j Joke
	if err := json.NewDecoder(resp.Body).Decode(&j); err != nil {
		return nil, err
	}
	return &j, nil
}

func saveJoke(j *Joke) error {
	_, err := db.Exec(
		`INSERT INTO jokes(category,type,joke,setup,delivery) VALUES($1,$2,$3,$4,$5)`,
		j.Category, j.Type, j.Joke, j.Setup, j.Delivery,
	)
	return err
}

func jokeHandler(w http.ResponseWriter, r *http.Request) {
	j, err := fetchJoke()
	if err != nil {
		http.Error(w, "fetch joke failed: "+err.Error(), 500)
		return
	}
	if err := saveJoke(j); err != nil {
		http.Error(w, "save joke failed: "+err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(j)
}

func memeHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://meme-api.com/gimme/programming_memes")
	if err != nil {
		http.Error(w, "fetch meme failed: "+err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	var m Meme
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		http.Error(w, "invalid meme JSON: "+err.Error(), 500)
		return
	}

	t := template.Must(template.New("meme").Parse(`
		<!DOCTYPE html>
		<html lang="en">
		<head><meta charset="UTF-8"><title>{{ .Title }}</title></head>
		<body style="text-align:center;font-family:sans-serif">
			<h1>{{ .Title }}</h1>
			<p><a href="{{ .PostURL }}">View on Reddit</a></p>
			<img src="{{ .URL }}" alt="{{ .Title }}" style="max-width:90%;height:auto"/>
		</body>
		</html>
	`))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, m)
}

func main() {
	setupDB()
	http.HandleFunc("/joke", jokeHandler)
	http.HandleFunc("/meme", memeHandler)
	addr := ":8080"
	log.Printf("Listening on http://localhost%s ...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
