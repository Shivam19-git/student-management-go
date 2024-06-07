package main

import (
	"fmt"

	"project/backend/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port :=8080
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is Home!")
	})

	// use handlers router function for all requests to /api/signin
	api := app.Group("/api")
	handler.SetupRoutes(api)

	fmt.Print("Server is listening on http://localhost:", port)
	app.Listen(":8080")
}
