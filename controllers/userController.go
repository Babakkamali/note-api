package controllers

import (
	"github.com/babakkamali/note-api/services"
	"github.com/babakkamali/note-api/utils"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (uc *UserController) AuthenticateOrRegister(c *fiber.Ctx) error {
    var input struct {
        PhoneNumber string `json:"phone_number"`
    }

    if err := c.BodyParser(&input); err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
    }

	// Validate the phone number with regex
	validPhoneNumber := utils.ValidatePhoneNumber(input.PhoneNumber)
	if !validPhoneNumber {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid phone number format")
	}

	_, err := uc.userService.GenerateAndSendSMSToken(input.PhoneNumber)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
    }

    return c.JSON(fiber.Map{"status": "success", "message": "SMS token sent"})

}

func (uc *UserController) ValidateSMSTokenAndLogin(c *fiber.Ctx) error {
    var input struct {
        PhoneNumber string `json:"phone_number"`
        SMSToken    string `json:"sms_token"`
    }

    if err := c.BodyParser(&input); err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
    }

    jwtToken, err := uc.userService.ValidateTokenAndLogin(input.PhoneNumber, input.SMSToken)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
    }

    return c.JSON(fiber.Map{"token": jwtToken})
}