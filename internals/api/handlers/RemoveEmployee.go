package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lillez7/HRMS/internals/database"
	"github.com/lillez7/HRMS/internals/types"
)

func RemoveEmployee(c *fiber.Ctx) error {
	var id int

	err := c.BodyParser(&id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	database.DB.Delete(&types.Employee{}, id)
	message := fmt.Sprintf("successfully removed user with id %d", id)
	return c.Status(200).JSON(fiber.Map{
		"data": message,
	})
}
