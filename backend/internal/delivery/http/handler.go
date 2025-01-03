package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/google/uuid"

	"course_photos/internal/models"
	"course_photos/internal/usecase"
)

type Config struct {
	Addr string

	SessionCookieDomain string `validate:"required"`
	SessionCookieTTL    int    `validate:"seconds"`
	CookieSameSite      string `validate:"required"`
	AllowLocalhost      bool
	SecureCookies       bool
}

type handler struct {
	cfg Config
	uc  *usecase.UseCase
	app *fiber.App
}

func (h *handler) Run() error {
	return h.app.Listen(h.cfg.Addr)
}

func (h *handler) bindRoutes() {
	h.app.Post("/register", h.handlerRegister)
	h.app.Post("/login", h.handlerLogin)
	h.app.Delete("/logout", h.handlerLogout)

	authorizedGroup := h.app.Group("/", h.authMiddleware)
	{
		meGroup := authorizedGroup.Group("/me")
		{
			meGroup.Get("/", h.handlerGetMe)
			meGroup.Get("/studios", h.handlerGetMyStudios)
			meGroup.Post("/studios", h.handlerCreateStudio)
			//meGroup.Patch("/studios/:id", h.handlerUpdateMyStudio)
			meGroup.Delete("/studios/:id", h.handlerDeleteMyStudio)
			meGroup.Get("/studios/:id/bookings", h.handlerGetMyStudioBookings)
			meGroup.Get("/bookings", h.handlerGetMyBookings)
			meGroup.Delete("/bookings/:id", h.handlerDeleteBooking)
		}
		authorizedGroup.Get("/studios", h.handlerGetStudios)
		authorizedGroup.Get("/studios/:id/available-hours", h.handlerGetStudioAvailableHours)
		authorizedGroup.Post("/studios/:id/bookings", h.handlerCreateBookings)
		authorizedGroup.Get(
			"/photo/:id",
			h.getCacheMiddleware(10*time.Second),
			h.handlerGetPhoto,
		)
		authorizedGroup.Post("/photo", h.handlerUploadPhoto)
	}

	for _, route := range h.app.GetRoutes() {
		fmt.Println(route.Method, route.Path)
	}
}

const (
	sessionCookie = "session"
	localsKeyUser = "user"
)

func (h *handler) authMiddleware(c *fiber.Ctx) error {
	sessionId, err := uuid.Parse(c.Cookies(sessionCookie))
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	user, err := h.uc.Auth(c.Context(), sessionId)
	if err != nil {
		h.clearCookie(c, sessionCookie)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	c.Locals(localsKeyUser, user)

	return c.Next()
}
func getUser(c *fiber.Ctx) models.User {
	return c.Locals(localsKeyUser).(models.User)
}
func (h *handler) getCacheMiddleware(expiration time.Duration) fiber.Handler {
	return cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("noCache") == "true"
		},
		Expiration: expiration,
		ExpirationGenerator: func(c *fiber.Ctx, cfg *cache.Config) time.Duration {
			if c.Response().StatusCode() == http.StatusNoContent {
				return 0
			}

			return cfg.Expiration
		},
		CacheControl: true,
	})
}
