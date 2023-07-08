package routes

import (
	usercontroller "st/cmd/api_server/controllers/user_controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes() *fiber.App {
	app := fiber.New()

	app.Get("/users", usercontroller.Index)

	return app
}
