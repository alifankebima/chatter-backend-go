package controllers

import (
	"chatter-backend-go/src/models"

	"github.com/gofiber/fiber/v2"
)

// func RegisterUser(c *fiber.Ctx) error {

// }

func GetProfile(c *fiber.Ctx) error {
	user := models.SelectUser(c.Params("id"))
	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	return c.JSON(user)
}
