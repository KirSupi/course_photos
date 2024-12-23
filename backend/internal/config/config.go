package config

import (
	"course_photos/internal/delivery/http"
	"course_photos/pkg/postgres"
)

type Config struct {
	Postgres postgres.Config
	HTTP     http.Config
}
