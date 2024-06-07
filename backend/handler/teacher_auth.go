package handler

import "github.com/gofiber/fiber/v2"

func TeacherSignIn(c *fiber.Ctx) error {
	return c.SendString("This is Teacher Signin")
}