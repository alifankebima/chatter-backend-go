package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	v1 := app.Group("/api/v1")

	User(v1.Group("/user"))
}
