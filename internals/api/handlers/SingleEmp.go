package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lillez7/HRMS/internals/database"
	"github.com/lillez7/HRMS/internals/types"
)

func GetEmpByID(c *fiber.Ctx) error {
	var id int

	err := c.BodyParser(&id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	var employee types.Employee

	database.DB.First(&employee, id)

	return c.Status(200).JSON(fiber.Map{
		"data": employee,
	})
}
