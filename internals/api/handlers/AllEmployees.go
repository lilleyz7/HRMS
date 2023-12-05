package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lillez7/HRMS/internals/database"
	"github.com/lillez7/HRMS/internals/types"
)

func GetAllEmployees(c *fiber.Ctx) error {
	var employees []types.Employee
	err := database.DB.Find(&employees)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": employees,
	})

}
