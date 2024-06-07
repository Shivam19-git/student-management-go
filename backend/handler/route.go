package handler

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app fiber.Router) error{
	signin := app.Group("/signin")

	// define sub routes
	signin.Get("/admin", AdminSignIn)
    signin.Get("/teacher", TeacherSignIn)
    signin.Get("/student", StudentSignIn)
	
    return nil
}