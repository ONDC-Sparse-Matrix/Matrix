package main

import (
	"regional_server/configs"
	"regional_server/routes" //add this
	"regional_server/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	configs.ConnectDB()
	controllers.InitCronJob();
	routes.MerchantRouter(app)
	routes.MapRouter(app)

	app.Listen(":5000")
}
