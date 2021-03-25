package session

import (
	"time"

	fsession "github.com/gofiber/fiber/v2/middleware/session"
)

var Store = fsession.New(fsession.Config{
	Expiration:     5 * time.Minute,
	CookieHTTPOnly: true,
})
