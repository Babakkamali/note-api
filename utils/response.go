package utils

import (
    "github.com/gofiber/fiber/v2"
)

type Response struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Meta    interface{} `json:"meta,omitempty"`
}

func SendResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
    return c.Status(status).JSON(&Response{
        Status:  "success",
        Message: message,
        Data:    data,
    })
}

func SendErrorResponse(c *fiber.Ctx, status int, message string) error {
    return c.Status(status).JSON(&Response{
        Status:  "error",
        Message: message,
    })
}