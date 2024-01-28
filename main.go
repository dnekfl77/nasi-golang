package main

import (
	"log"
	"nasi-golang/database"
	"nasi-golang/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	database.Connect()

	engine := django.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main-layout",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"message": "Hello, World!",
		})
	})

	// Set routers
	router.InstallRouter(app)

	log.Fatal(app.Listen(":3000"))
}
