package http

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"

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

func (h *handler) handlerGetStudios(c *fiber.Ctx) error {
	data, err := h.uc.GetStudios(c.Context())
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

	data, err := h.uc.GetStudioAvailableHours(c.Context(), studioId, date)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(data)
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
func (h *handler) handlerDeleteStudio(c *fiber.Ctx) error {
	studioId, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = h.uc.DeleteStudio(c.Context(), getUser(c).Id, int64(studioId))
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}

//func (h *handler) handlerUpdateStudio(c *fiber.Ctx) error {
//
//}
//
//func (h *handler) handlerDeleteStudio(c *fiber.Ctx) error {
//
//}
//
//func (h *handler) handlerGetStudioBookings(c *fiber.Ctx) error {
//
//}
//
//func (h *handler) handlerBookStudio(c *fiber.Ctx) error {
//
//}
