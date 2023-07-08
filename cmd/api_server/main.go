package main

import (
	"st/cmd/api_server/routes"
	"st/pkg"
	"st/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	if err := pkg.LoadENV(".env"); err != nil {
		panic(err)
	}

	config.InitAuthConfig()

	pkg.InitEnt()

	app := fiber.New()
	router := app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowCredentials: true,
	}))


	api := router.Group("/api")
	v1 := api.Group("/v1")
	
	v1.Mount("",routes.AuthRoutes())
	v1.Mount("", routes.UserRoutes())
	
	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
