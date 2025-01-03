package http

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"course_photos/internal/usecase"
	"course_photos/pkg/dates"
)

func (h *handler) handlerRegister(c *fiber.Ctx) error {
	req := usecase.RegisterRequest{}

	err := c.BodyParser(&req)
	if err != nil {
		return fiber.ErrBadRequest
	}

	sessionId, err := h.uc.Register(c.Context(), req)
	if err != nil {
		return err
	}

	h.setCookie(c, sessionCookie, sessionId.String(), nil, true)

	return c.SendStatus(http.StatusOK)
}
func (h *handler) handlerLogin(c *fiber.Ctx) error {
	req := usecase.LoginRequest{}

	err := c.BodyParser(&req)
	if err != nil {
		return fiber.ErrBadRequest
	}

	sessionId, err := h.uc.Login(c.Context(), req)
	if err != nil {
		return err
	}

	h.setCookie(c, sessionCookie, sessionId.String(), nil, true)

	return c.SendStatus(http.StatusOK)
}

func (h *handler) handlerLogout(c *fiber.Ctx) error {
	sessionIdStr := c.Cookies(sessionCookie)
	if sessionIdStr == "" {
		return c.SendStatus(http.StatusOK)
	}

	sessionId, err := uuid.Parse(sessionIdStr)
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = h.uc.Logout(c.Context(), sessionId)
	if err != nil {
		return err
	}

	h.clearCookie(c, sessionCookie)

	return c.SendStatus(http.StatusOK)
}

func (h *handler) handlerGetStudios(c *fiber.Ctx) error {
	data, err := h.uc.GetStudios(c.Context(), getUser(c).Id)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(data)
}
func (h *handler) handlerGetStudioAvailableHours(c *fiber.Ctx) error {
	studioId, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	date := dates.Date{}

	err = date.UnmarshalText([]byte(c.Query("date")))
	if err != nil {
		return fiber.ErrBadRequest
	}

	data, err := h.uc.GetStudioAvailableHours(c.Context(), int64(studioId), date)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(data)
}
func (h *handler) handlerCreateBookings(c *fiber.Ctx) error {
	req := usecase.CreateBookingsRequest{}

	err := c.ParamsParser(&req)
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = c.BodyParser(&req)
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = h.uc.CreateBookings(c.Context(), getUser(c).Id, req)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}
func (h *handler) handlerGetPhoto(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	data, err := h.uc.GetPhoto(c.Context(), int64(id))
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return c.SendStatus(fiber.StatusNoContent)
	}

	mimeType := http.DetectContentType(data)
	c.Set(fiber.HeaderContentType, mimeType)

	_, err = c.Write(data)
	if err != nil {
		return err
	}

	c.Status(http.StatusOK)

	return nil
}
func (h *handler) handlerGetMe(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(getUser(c))
}
func (h *handler) handlerGetMyStudios(c *fiber.Ctx) error {
	data, err := h.uc.GetMyStudios(c.Context(), getUser(c).Id)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(data)
}
func (h *handler) handlerGetMyBookings(c *fiber.Ctx) error {
	data, err := h.uc.GetMyBookings(c.Context(), getUser(c).Id)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(data)
}
func (h *handler) handlerDeleteBooking(c *fiber.Ctx) error {
	bookingId, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = h.uc.DeleteMyBooking(c.Context(), getUser(c).Id, int64(bookingId))
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}

func (h *handler) handlerUploadPhoto(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return fiber.ErrBadRequest
	}

	f, err := file.Open()
	if err != nil {
		return fiber.ErrBadRequest
	}

	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	if !strings.HasSuffix(strings.ToLower(file.Filename), ".jpg") {
		return fiber.ErrBadRequest
	}

	photo, err := io.ReadAll(f)
	if err != nil {
		return fiber.ErrBadRequest
	}

	id, err := h.uc.UploadPhoto(c.Context(), getUser(c).Id, photo)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(id)
}
func (h *handler) handlerCreateStudio(c *fiber.Ctx) error {
	req := usecase.CreateStudioRequest{}

	err := c.BodyParser(&req)
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = h.uc.CreateStudio(c.Context(), getUser(c).Id, req)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}
func (h *handler) handlerDeleteMyStudio(c *fiber.Ctx) error {
	studioId, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = h.uc.DeleteMyStudio(c.Context(), getUser(c).Id, int64(studioId))
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}
func (h *handler) handlerGetMyStudioBookings(c *fiber.Ctx) error {
	studioId, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	data, err := h.uc.GetMyStudioBookings(c.Context(), getUser(c).Id, int64(studioId))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(data)
}
