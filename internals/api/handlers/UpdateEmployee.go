package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lillez7/HRMS/internals/database"
	"github.com/lillez7/HRMS/internals/types"
)

func UpdateUserInfo(c *fiber.Ctx) error {
	var body struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
		Salary    int    `json:"salary"`
		Id        int    `json:"id"`
	}

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid body",
		})
	}

	var employee types.Employee
	database.DB.First(&employee, body.Id)

	database.DB.Model(&employee).Updates(types.Employee{
		Firstname: body.FirstName,
		Lastname:  body.LastName,
		Email:     body.Email,
		Salary:    body.Salary,
	})

	return c.Status(200).JSON(fiber.Map{
		"data": employee,
	})

}

func UpdateTimeOff() {

}
