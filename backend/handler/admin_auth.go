package handler

import (
	"github.com/gofiber/fiber/v2"
)

func AdminSignIn(c *fiber.Ctx) error {
	return c.SendString("This is admin Signin")
}
