package router

import (
	"github.com/Chanchit1516/middleware"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// Middleware
	app.Post("/login", middleware.Login)
	app.Get("/", middleware.Accessible)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey:   []byte("secret"),
		ErrorHandler: func(c *fiber.Ctx, e error) error { return c.SendStatus(fiber.StatusUnauthorized) },
	}))

	app.Get("/restricted", middleware.Restricted)
}
