package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
)

////go:embed views
// var viewsAsssets embed.FS

func main()  {
	engine := handlebars.New("views", ".hbs")
	
	
	app := fiber.New(fiber.Config{
		Views: engine,
	})


	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", nil, "layouts/auth")
	})

	app.Get("/",func(c *fiber.Ctx) error {
		return c.Render("index", nil, "layouts/main")
	})

	if err := app.Listen(":8000"); err != nil  {
		panic(err)
	}
}
