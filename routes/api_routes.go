package routes

import (
	"SoundGadget-API/controllers"

	"github.com/gofiber/fiber/v2"
)

// routes to endpoints in this API

func ApiRoute(app *fiber.App) {

	app.Post("/soundgadget/", controllers.AddProduct)
	app.Get("/soundgadget/:name", controllers.GetProductByName)
	app.Get("/soundgadget/TypePrice/*", controllers.GetProductByPriceandType)
	app.Get("/soundgadget/TypeBrand/*", controllers.GetProductByTypeAndBrand)
	app.Put("/soundgadget/edit/:prtoedit", controllers.EditProductByName)
	app.Delete("/soundgadget/delete/:prtodeleted", controllers.DeleteProductByName)

}
