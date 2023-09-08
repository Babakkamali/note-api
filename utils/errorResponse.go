package utils

import (
    "github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
    Status  int    `json:"status"`
    Message string `json:"message"`
}

func SendErrorResponse(c *fiber.Ctx, status int, message string) error {
    return c.Status(status).JSON(&ErrorResponse{
        Status:  status,
        Message: message,
    })
}