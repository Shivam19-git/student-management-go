package handler

import "github.com/gofiber/fiber/v2"

func StudentSignIn(c *fiber.Ctx) error {
	return c.SendString("This is Student Signin")
}