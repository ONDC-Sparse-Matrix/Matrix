package main

import (
	"centralCDN/pkg/configs"
	"centralCDN/pkg/types"
	"centralCDN/pkg/utils"
	"encoding/json"

	// "encoding/json"
	"fmt"
	// "io/ioutil"
	"log"
	"strconv"
	"time"

	"context"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Initialise() {
	utils.InitCacheServerList()
	utils.InitPincode()
	utils.InitServerRangeList()
}

type Merchant struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PincodeMap struct {
	Pincode     string               `bson:"_id"`
	MerchantIDs []primitive.ObjectID `bson:"merchant_ids"`
}

var client *mongo.Client

func main() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(configs.EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to MongoDB")
	}
	defer client.Disconnect(ctx)

	config := fiber.Config{
		ServerHeader: "Cache Server",
		Prefork:      true,
		// Concurrency:  1024 * 512,
	}
	app := fiber.New(config)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: configs.EnvFrontendURL(),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	queueConnection := utils.ConnectQueue()
	defer queueConnection.Close()

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

	app.Post("/pincode/:pincode?", func(c *fiber.Ctx) error {

		var clientId types.Session
		pincode := c.Params("pincode")
		pincodeInt, _ := strconv.Atoi(pincode)
		go utils.UpdateFreqMap(pincodeInt)

		if err := c.BodyParser(&clientId); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}
		fmt.Println("Client ID: ", clientId.ClientID)

		fmt.Println("Pincode", pincode)
		body := utils.FetchMerchantData(pincode, clientId.ClientID)

		return c.SendString(body)

	})

	app.Post("/cache/:clientId", func(c *fiber.Ctx) error {
		clientId := c.Params("clientId")
		var response types.CachePayload
		err := json.Unmarshal(c.Body(), &response)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}
		var cacheResponse []types.PincodeInfo
		var clientCacheResponse []types.PincodeInfo
		cacheResponse = response.CacheResponse
		clientCacheResponse = response.ClientCacheResponse
		log.Println("Client ID: ", clientId)
		log.Println("Cache Response: ", len(cacheResponse))
		log.Println("Client Cache Response: ", len(clientCacheResponse))

		// //TODO: @DAGGER store the cacheResponse in the cache

		err = utils.PublishCache(queueConnection, clientId, cacheResponse)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to store cache response"})
		}
		return c.SendString("Stored Cache Response")
	})

	app.Get("/sse/:clientId",websocket.New(func(c *websocket.Conn){
		clientId := c.Params("clientId")
		utils.ConsumeCache(queueConnection, clientId, c)
	
	}))

	app.Get("/upload", websocket.New(func(c *websocket.Conn) {
		var (
			mt  int
			msg []byte
			err error
		)

		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}

			if string(msg) == "done" {
				break
			}

			err := utils.PublishSeeding(queueConnection, msg)

			if err != nil {
				log.Fatalf("Failed to publish a message: %v", err)
			}

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))

	//handling seeding events
	utils.ConsumeSeeding(queueConnection, client)

	app.Listen(":3001")

}
