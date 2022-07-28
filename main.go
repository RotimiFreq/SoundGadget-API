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
  
    // app.Get("/", func(c *fiber.Ctx) error {
    //     return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
    // })
  
    app.Listen(":8000")
}