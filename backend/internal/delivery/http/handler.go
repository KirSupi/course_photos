package http

import (
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	app fiber.App
}

func (h *handler) Run() error {
	h.bindRoutes()

	return h.app.Listen(":3000")
}

func (h *handler) bindRoutes() {
	h.app.Post("/register", h.handlerRegister)
	h.app.Post("/login", h.handlerLogin)

	h.app.Post("/studios", h.handlerCreateStudio)
	h.app.Put("/studios/:id", h.handlerUpdateStudio)
	h.app.Delete("/studios/:id", h.handlerDeleteStudio)
	h.app.Get("/studios/:id/bookings", h.handlerGetStudioBookings)
	h.app.Post("/studios/:id/book", h.handlerBookStudio)
}
