package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func LoginUser(c *fiber.Ctx) error {
	// logic for user login here
	
	phoneNumber := c.FormValue("phone_number")

	return c.JSON(fiber.Map{
		"message": "User logged in successfully",
		"user":    phoneNumber,
	})
}