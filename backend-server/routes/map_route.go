package routes

import (
	"regional_server/controllers"

	"github.com/gofiber/fiber/v2"
)

func MapRouter(app *fiber.App) {
	app.Get("/map/update/:pincode", controllers.UpdateMap)
}
