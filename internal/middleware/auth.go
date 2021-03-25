package middleware

import (
	"github.com/dmitryt/tiny-url-service-backend/internal/session"
	"github.com/gofiber/fiber/v2"
)

var publicRoutes = []string{"/api/v1/auth/login", "/api/v1/auth/register"}

func Auth(c *fiber.Ctx) error {
	for _, route := range publicRoutes {
		if route == c.Path() {
			return c.Next()
		}
	}
	store, err := session.Store.Get(c)
	if err != nil {
		panic(err)
	}

	uid := store.Get("uid")
	if uid == nil {
		return fiber.ErrUnauthorized
	}

	return c.Next()
}
