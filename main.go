package main

import (
	"github.com/Chanchit1516/router"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
