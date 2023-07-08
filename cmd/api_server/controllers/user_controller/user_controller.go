package usercontroller

import (
	"st/pkg"

	"github.com/gofiber/fiber/v2"
)


func Index(c *fiber.Ctx) error {
	users_, err := pkg.EntClient().User.Query().Limit(10).All(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"users": users_,
	})
}
