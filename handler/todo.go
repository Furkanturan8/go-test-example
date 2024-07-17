package handler

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"todoAppWithTest/model"
	"todoAppWithTest/service"
)

func CreateTodoHandler(c *fiber.Ctx) error {
	todo := new(model.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	createdTodo := service.CreateTodoService(*todo)
	return c.Status(fiber.StatusCreated).JSON(createdTodo)
}

func GetTodosHandler(c *fiber.Ctx) error {
	todos := service.GetTodosService()
	return c.JSON(todos)
}

func UpdateTodoHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	todo := new(model.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	updatedTodo, err := service.UpdateTodoService(id, *todo)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(updatedTodo)
}

func DeleteTodoHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	err = service.DeleteTodoService(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
