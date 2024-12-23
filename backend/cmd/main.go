package main

import (
	"log"

	"github.com/jmoiron/sqlx"

	"course_photos/internal/config"
	"course_photos/internal/delivery/http"
	"course_photos/internal/repository"
	"course_photos/internal/usecase"
	"course_photos/pkg/postgres"
)

func main() {
	cfg := config.Config{}

	err := config.Load(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.Connect(cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)

	r := repository.New(db)
	uc := usecase.New(r)

	server := http.New(cfg.HTTP, uc)

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
