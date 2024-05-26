package main

import (
	"regional_server/configs"
	"regional_server/routes" //add this

	controllers "command-line-arguments/home/death/Desktop/linux_backup/sdsLabs/HACKATHONS/Matrix/Matrix/backendServer/controllers/merchants_controller.go"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()
	controllers.InitCronJob();
	routes.MerchantRouter(app)
	routes.MapRouter(app)

	app.Listen(":5000")
}
