package main

import (
	"centralCDN/pkg/utils"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Initialise() {
	utils.InitCacheServerList()
	utils.InitPincode()
	utils.InitServerRangeList()
}

func main() {
	config := fiber.Config{
		ServerHeader: "Cache Server",
		Prefork:      true,
		// Concurrency:  1024 * 512,
	}
	app := fiber.New(config)
	Initialise()
	var count int
	app.Get("/", func(c *fiber.Ctx) error {
		count++
		fmt.Println("Welcome ", count)

		time.Sleep(2 * time.Second)
		fmt.Println("Endd ", count)
		count--
		return c.SendString("Hello, World!")
	})

	app.Get("/pincode/:pincode?", func(c *fiber.Ctx) error {

		pincode := c.Params("pincode")

		fmt.Println("Pincode", pincode)
		body := utils.FetchMerchantData(pincode)
		// pincode := c.Params("pincode")
		// jsonData := []byte(`[
		// 	{"pincode" : "22114097", "merchantList" : ["merchant1", "merchant2", "merchant3"]},
		// 	{"pincode" : "22114098", "merchantList" : ["merchant1", "merchant2", "merchant3"]},
		// 	{"pincode" : "22114099", "merchantList" : ["merchant1", "merchant2", "merchant3"]},
		// 	{"pincode" : "22114096", "merchantList" : ["merchant1", "merchant2", "merchant3"]},
		// 	{"pincode" : "22114095", "merchantList" : ["merchant1", "merchant2", "merchant3"]},
		// 	{"pincode" : "22114094", "merchantList" : ["merchant1", "merchant2", "merchant3"]},
		// 	{"pincode" : "22114093", "merchantList" : ["merchant1", "merchant2", "merchant3"]}
		// ]`)

		// go utils.FetchMerchantData(pincode)
		return c.SendString(body)
	})

	// type MerchantAddRequestBody struct {
	// 	Pincode      string   `json:"pincode"`
	// 	MerchantList []string `json:"merchantList"`
	// }
	// var merchantAddRequestBody MerchantAddRequestBody
	// app.Post("/merchant/add", func(c *fiber.Ctx) error {
	// 	if err := c.BodyParser(&merchantAddRequestBody); err != nil {
	// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	// 	}
	// 	merchantList := merchantAddRequestBody.MerchantList
	// 	baseUrl := fmt.Sprintf("http://%s:%s/merchant/new?pincode=%s&date=%s", types.Server.Host, types.Server.Port, pincode, lastModified.String())

	// 	resp, _ := http.Get(baseUrl)
	// 	body, _ := ioutil.ReadAll(resp.Body) // {status: 200}
	// 	if string(body) != "200" {
	// 		types.PincodeInfoList = cacheMiss(pincode, c)
	// 	}
	// 	return c.SendString("Recieved Response")
	// })

	app.Listen(":3001")
}

// /merchant/pincode/:pincode
