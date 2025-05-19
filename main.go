package main

import (
	"crud_api/config"
	"crud_api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectDB()
	routes.SetupRoutes(app)

	app.Listen(":3000")
}
