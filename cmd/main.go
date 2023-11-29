package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lillez7/HRMS/internals/api/handlers"
)

func main() {
	app := fiber.New()

	app.Get("/", handlers.GetAllEmployees)
}
