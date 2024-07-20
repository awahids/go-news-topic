package main

import (
	"log"
	"net/http"
	"news-topic-api/internal/db"
	"news-topic-api/internal/routes"
	"time"
)

func main() {
	// init db
	db.NewPostgresDB()

	// init routes
	r := routes.InitRoutes()

	server := &http.Server{
		Addr:           ":9000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Server listening on %s", server.Addr)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
