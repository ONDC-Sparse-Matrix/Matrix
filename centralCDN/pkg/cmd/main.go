package main

import (
	"centralCDN/pkg/types"
	"centralCDN/pkg/utils"
	"fmt"
	"log"
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
	client,err = mongo.Connect( ctx,options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
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
		AllowOrigins: "http://localhost:3000", //!.env frontend url
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

	app.Get("/pincode/:pincode?", func(c *fiber.Ctx) error {

		pincode := c.Params("pincode")

		fmt.Println("Pincode", pincode)
		body := utils.FetchMerchantData(pincode)

		return c.SendString(body)
	})

	app.Post("/cache/:clientId", func(c *fiber.Ctx) error {
		clientId := c.Params("clientId")
		var response types.CachePayload

		if err := c.BodyParser(&response); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		cacheResponse := response.CacheResponse
		clientCacheResponse := response.ClientCacheResponse
		println("Cache Response: ", clientCacheResponse)
		//TODO: @DAGGER store the cacheResponse in the cache

		err := utils.PublishCache(queueConnection, clientId, cacheResponse)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to store cache response"})
		}
		return c.SendString("Stored Cache Response")
	})

	app.Get("/sse/:clientId", func(c *fiber.Ctx) error {
		clientId := c.Params("clientId")
		c.Set("Content-Type", "text/event-stream")
		c.Set("Connection", "keep-alive")

		msgs, err := utils.ConsumeCache(queueConnection, clientId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to consume cache response"})
		}

		for msg := range msgs {
			c.Write([]byte("data: " + string(msg.Body) + "\n\n"))
		}
		return nil
	})

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
	go utils.HandleSeedQueue(queueConnection, client)

	app.Listen(":3001")

}


