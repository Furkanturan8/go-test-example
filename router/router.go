package router

import (
	"github.com/gofiber/fiber/v2"
	"todoAppWithTest/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/todos", handler.CreateTodoHandler)
	api.Get("/todos", handler.GetTodosHandler)
	api.Put("/todos/:id", handler.UpdateTodoHandler)
	api.Delete("/todos/:id", handler.DeleteTodoHandler)
}
