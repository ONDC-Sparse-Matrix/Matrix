package main

import (
	"centralCDN/pkg/utils"
	"fmt"
	"log"
	"time"

	"context"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	// "fmt"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	amqp "github.com/rabbitmq/amqp091-go"
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
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017")) //!.env
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
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
		AllowOrigins: "http://localhost:3000", //!.env
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

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
		var cacheResponse string
		var clientCacheResponse string

		if err := c.BodyParser(&cacheResponse); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if err := c.BodyParser(&clientCacheResponse); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		//TODO: @DAGGER store the cacheResponse in the cache

		cachePubChannel, err := conn.Channel()
		if err != nil {
			log.Fatalf("Failed to open a channel: %v", err)
		}
		defer cachePubChannel.Close()

		cacheQueueName := clientId
		_, err = cachePubChannel.QueueDeclare(cacheQueueName, true, false, false, false, nil)
		if err != nil {
			log.Fatalf("Failed to declare a queue: %v", err)
		}
		cache,err := json.Marshal(cacheResponse)
		if err != nil {
			log.Fatalf("Failed to marshal cache response: %v", err)
		}
		err = cachePubChannel.Publish("", cacheQueueName, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(cache),
		})
		if err != nil {
			log.Fatalf("Failed to publish a message: %v", err)
		}

		return c.SendString("Stored Cache Response")
	})

	app.Get("/sse/:clientId", func(c *fiber.Ctx) error {
		clientId := c.Params("clientId")
		c.Set("Content-Type", "text/event-stream")
		c.Set("Connection", "keep-alive")


		cacheSubChannel, err := conn.Channel()
		if err != nil {
			log.Fatalf("Failed to open a channel: %v", err)
		}
		defer cacheSubChannel.Close()

		cacheQueueName := clientId
		q, err := cacheSubChannel.QueueDeclare(cacheQueueName, true, false, false, false, nil)
		if err != nil {
			log.Fatalf("Failed to declare a queue: %v", err)
		}

		msgs, err := cacheSubChannel.Consume(q.Name, "", true, false, false, false, nil)
		if err != nil {
			log.Fatalf("Failed to register a consumer: %v", err)
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

		type msgObj struct {
			Name     string  `json:"name"`
			Email    string  `json:"email"`
			Pincodes []int32 `json:"pincodes"`
		}

		pubChannel, err := conn.Channel()
		if err != nil {
			log.Fatalf("Failed to open a channel: %v", err)
		}
		defer pubChannel.Close()

		queuename := "merchant_data"
		_, err = pubChannel.QueueDeclare(queuename, true, false, false, false, nil)
		if err != nil {
			log.Fatalf("Failed to declare a queue: %v", err)
		}

		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}

			if string(msg) == "done" {
				break
			}

			var data msgObj
			if err := json.Unmarshal(msg, &data); err != nil {
				log.Println("json unmarshal error:", err)
				break
			}

			err = pubChannel.Publish("", queuename, false, false, amqp.Publishing{
				ContentType: "text/plain",
				Body:        msg,
			})
			if err != nil {
				log.Fatalf("Failed to publish a message: %v", err)
			}

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))

	consChannel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer consChannel.Close()

	queuename := "merchant_data"
	q, err := consChannel.QueueDeclare(queuename, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := consChannel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	type msgObj struct {
		Name     string  `json:"name"`
		Email    string  `json:"email"`
		Pincodes []int32 `json:"pincodes"`
	}
	go func() {
		for msg := range msgs {
			var data msgObj
			if err := json.Unmarshal(msg.Body, &data); err != nil {
				log.Println("json unmarshal error:", err)
				break
			}
			merchantCollection := client.Database("matrix_merchants").Collection("merchant_details")
			merchant := Merchant{
				Name:  data.Name,
				Email: data.Email,
			}
			result, err := merchantCollection.InsertOne(context.Background(), merchant)
			if err != nil {
				break
			}
			merchantID := result.InsertedID.(primitive.ObjectID)

			for _, pincode := range data.Pincodes {
				databaseName := getDatabaseName(pincode)
				if databaseName == "0" {
					break
				}
				pincodeCollection := client.Database(databaseName).Collection("pincode_map")
				filter := bson.M{"_id": pincode}
				update := bson.M{
					"$addToSet": bson.M{
						"merchant_ids": merchantID,
					},
				}
				_, err := pincodeCollection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
				if err != nil {
					break
				}
			}
		}

	}()

	app.Listen(":3001")

}

func getDatabaseName(pincode int32) string {
	regionCode := pincode / 100000
	switch regionCode {
	case 1, 2:
		return "matrix_map_1"
	case 3, 4:
		return "matrix_map_2"
	case 5, 6:
		return "matrix_map_3"
	case 7, 8, 9:
		return "matrix_map_4"
	default:
		return "0"
	}
}
