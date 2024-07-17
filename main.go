package main

import (
	"github.com/gofiber/fiber/v2"
	"todoAppWithTest/router"
)

func main() {
	app := fiber.New()
	router.SetupRoutes(app)
	app.Listen(":3000")
}
