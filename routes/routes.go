package routes

import (
	"crud_api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/mahasiswa", handlers.GetMahasiswa)
	api.Post("/mahasiswa", handlers.CreateMahasiswa)
	api.Put("/mahasiswa/:id", handlers.UpdateMahasiswa)
	api.Delete("/mahasiswa/:id", handlers.DeleteMahasiswa)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the CRUD API")
	})
}
