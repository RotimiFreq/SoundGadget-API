package main

import (
	"SoundGadget-API/configs"

	"SoundGadget-API/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	routes.ApiRoute(app)

	app.Listen(":8080")
}
