package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/lillez7/HRMS/internals/api/handlers"
	"github.com/lillez7/HRMS/internals/migrations"
)

func init() {
	Initialize()
	migrations.AutoMigrate()
}

func main() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(cors.New())

	port := os.Getenv("PORT")

	app.Get("/", handlers.GetAllEmployees)
	app.Get("/single", handlers.GetEmpByID)
	app.Post("/new", handlers.AddEmployee)
	app.Delete("/delete", handlers.RemoveEmployee)
	app.Put("/update", handlers.UpdateUserInfo)

	app.Listen(":" + port)
}
