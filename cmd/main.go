package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/lillez7/HRMS/internals/api/handlers"
	"github.com/lillez7/HRMS/internals/database"
)

func init() {
	LoadDotEnv()
	database.ConnectToDB()
}

func main() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(cors.New())

	port := os.Getenv("PORT")

	app.Get("/", handlers.GetAllEmployees)

	app.Listen(":" + port)
}
