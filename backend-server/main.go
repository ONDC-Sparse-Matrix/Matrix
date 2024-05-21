package main

import (
	"regional_server/configs"
	"regional_server/routes" //add this

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	routes.MerchantRouter(app)
	routes.MapRouter(app)

	app.Listen(":5000")
}
