package routes

import (
	"os"
	"st/cmd/api_server/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AuthRoutes() *fiber.App {
	app := fiber.New()

	app.Post("/login", func(c *fiber.Ctx) error {
		type LoginInput struct {
			Identity string `json:"identity"`
			Password string `json:"password"`
		}

		var input LoginInput
		if err := c.BodyParser(&input); err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		identity := input.Identity
		pass := input.Password
		if identity != "admin" || pass != "password" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["identity"] = identity
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"token": t, "env": os.Getenv("APP_ENV")})
	})

	protected := app.Group("", middleware.Protected())
	protected.Get("/me", func(c *fiber.Ctx) error {
		return nil
	})
	protected.Post("/logout", func(c *fiber.Ctx) error {
		return nil
	})

	return app
}
