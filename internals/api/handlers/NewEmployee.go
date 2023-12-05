package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lillez7/HRMS/internals/database"
	"github.com/lillez7/HRMS/internals/types"
)

func AddEmployee(c *fiber.Ctx) error {
	var employee types.Employee

	err := c.BodyParser(&employee)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	result := database.DB.Create(&employee)

	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": result.Error,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": employee.ID,
	})
}
