package routes

import (
	"regional_server/controllers"

	"github.com/gofiber/fiber/v2"
)

func MerchantRouter(app *fiber.App) {
	app.Get("/merchants/:pincode", controllers.GetMerchants)
	app.Post("/merchants/new",controllers.AddMerchants)
	// app.Put("/merchants/update/:id",controllers.UpdateMerchant)
}
