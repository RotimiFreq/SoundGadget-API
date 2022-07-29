package routes

import(
	"github.com/gofiber/fiber/v2"
	"SoundGadget-API/controllers"
)


func ApiRoute(app *fiber.App){


	app.Put("/mixers", controllers.CreateMixer)
	app.Get("/mixers/:Id", controllers.GetMixer)




	
}