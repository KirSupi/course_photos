package http

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

const localhost = "localhost"

func (h *handler) setCookie(
	c *fiber.Ctx,
	name, value string,
	expiredAt *time.Time,
	httpOnly ...bool,
) {
	cookie := &fiber.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Secure:   h.cfg.SecureCookies,
		Domain:   h.cfg.SessionCookieDomain,
		HTTPOnly: true,
		SameSite: h.cfg.CookieSameSite,
	}

	if h.cfg.AllowLocalhost && strings.HasPrefix(c.Hostname(), localhost) {
		cookie.Domain = localhost
		//cookie.SameSite = h.cfg.CookieSameSite
	}

	if expiredAt != nil {
		cookie.Expires = *expiredAt
	}

	if len(httpOnly) != 0 {
		cookie.HTTPOnly = httpOnly[0]
	}

	c.Cookie(cookie)
}

func (h *handler) clearCookie(c *fiber.Ctx, name string) {
	cookie := &fiber.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		Secure:   h.cfg.SecureCookies,
		Domain:   h.cfg.SessionCookieDomain,
		HTTPOnly: true,
		Expires:  time.Unix(0, 0),
		SameSite: h.cfg.CookieSameSite,
	}

	if h.cfg.AllowLocalhost && strings.HasPrefix(c.Hostname(), localhost) {
		cookie.Domain = localhost
		//cookie.SameSite = "None"
	}

	c.Cookie(cookie)
}
