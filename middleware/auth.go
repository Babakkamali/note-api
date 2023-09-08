package middleware

import (
    "github.com/babakkamali/note-api/utils"
    "github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
    token := c.Get("Authorization")
    
    // Strip the "Bearer " prefix if it exists
    if len(token) > 7 && token[0:7] == "Bearer " {
        token = token[7:]
    }

    if token == "" {
        return utils.SendErrorResponse(c, fiber.StatusUnauthorized, "Missing authentication token")
    }

    claims, err := utils.ParseJWT(token)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusUnauthorized, "Invalid token")
    }

    c.Locals("userID", claims.UserID)
    return c.Next()
}