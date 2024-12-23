package http

import (
	"github.com/gofiber/fiber/v2"

	"course_photos/internal/usecase"
)

type Server interface {
	Run() error
}

func New(cfg Config, uc *usecase.UseCase) Server {
	h := &handler{
		cfg: cfg,
		uc:  uc,
		app: fiber.New(fiber.Config{
			DisableStartupMessage: true,
		}),
	}

	h.bindRoutes()

	return h
}
