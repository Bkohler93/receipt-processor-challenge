package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pressly/goose"

	_ "github.com/lib/pq"
)

const port = ":8080"

func main() {
	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cfg := newConfig(db)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/receipts/process", cfg.handlePostReceiptsProcess)
	r.Get("/receipts/{id}/points", cfg.handleGetReceiptsPoints)

	fmt.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func setupDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=dev_user password=dev_pass host=localhost port=5432 dbname=dev_db sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("unable to open db - %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("unable to ping db - %v", err)
	}

	goose.SetDialect("postgres")
	if err := goose.Up(db, "sql/migrations"); err != nil {
		return nil, fmt.Errorf("failed to run migrations - %v", err)
	}
	return db, nil
}
