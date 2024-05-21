package main

import (
	"cache-server/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
	// "net/http"
	"time"
)

func Initialize() {
	utils.InitServer()
}

func main() {
	Initialize()
	cache1 := cache.New(48*time.Hour, 48*time.Hour)
	config := fiber.Config{ServerHeader: "Cache Server", Prefork: true}
	app := fiber.New(config)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/pincode/:pincode?", func(c *fiber.Ctx) error {
		pincode := c.Params("pincode")
		fmt.Println("pincode", pincode)
		jsonData := utils.CheckPincode(pincode, c, cache1)
		return c.SendString(string(jsonData))
	})
	type UpdateRequestBody struct {
		PincodeList []string `json:"pincodeList"`
	}

	app.Post("/update/", func(c *fiber.Ctx) error {
		// Parse the request body into a struct
		var requestBody UpdateRequestBody
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}
		utils.CheckCache(requestBody.PincodeList, cache1)
		// Process the data as needed
		// response := fiber.Map{
		// 	"message": "Received POST request",
		// 	"name":    requestBody.Name,
		// 	"email":   requestBody.Email,
		// }

		// Send a JSON response
		return c.SendString("Recieved Response")
	})

	// app.Get

	app.Listen(":4000")
}
