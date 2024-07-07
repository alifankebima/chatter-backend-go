package routes

import (
	"chatter-backend-go/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func User(user fiber.Router) {
	user.Get("/:id", controllers.GetProfile)
}
