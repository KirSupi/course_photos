package http

import (
	"github.com/gofiber/fiber/v2"
)

type Server interface {
	Run() error
}

func New() Server {
	return &handler{
		app: fiber.New(),
	}
}
