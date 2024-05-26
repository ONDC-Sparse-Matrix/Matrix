package utils

import (
	"centralCDN/pkg/configs"
	"centralCDN/pkg/types"
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/contrib/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectQueue() *amqp.Connection {
	conn, err := amqp.Dial(configs.EnvRabbitMQURI())
	if err != nil {
		panic(err)
	}
	return conn
}

func getPubChannel(Connection *amqp.Connection, channelName string) *amqp.Channel {
	ch, err := Connection.Channel()
	if err != nil {
		panic(err)
	}

	ch.QueueDeclare(channelName, true, false, false, false, nil)
	return ch
}

func getSubChannel(Connection *amqp.Connection, channelName string) *amqp.Channel {
	ch, err := Connection.Channel()
	if err != nil {
		panic(err)
	}

	ch.QueueDeclare(channelName, true, false, false, false, nil)
	return ch
}

func PublishCache(Connection *amqp.Connection, clientId string, cache []types.PincodeInfo) error {
	ch := getPubChannel(Connection, clientId)
	defer ch.Close()
	body, err := json.Marshal(cache)
	if err != nil {
		return err
	}
	ch.Publish("", clientId, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	fmt.Println("Published")
	return nil
}

func ConsumeCache(Connection *amqp.Connection, clientId string, c *websocket.Conn) {
	ch := getSubChannel(Connection, clientId)
	msgs, err := ch.Consume(clientId, "", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
	}

	for msg := range msgs {
		fmt.Println("Message Received", string(msg.Body))
		err := c.WriteMessage(websocket.TextMessage, msg.Body)
		if err != nil {
			fmt.Println("Error sending SSE:", err)
		}
	}

}

func PublishSeeding(Connection *amqp.Connection, data []byte) error {
	ch := getPubChannel(Connection, "merchant_data")
	defer ch.Close()
	ch.Publish("", "merchant_data", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        data,
	})
	return nil
}

func ConsumeSeeding(Connection *amqp.Connection, client *mongo.Client) {
	ch := getSubChannel(Connection, "merchant_data")
	msgs, err := ch.Consume("merchant_data", "", true, false, false, false, nil)
	if err != nil {
		return
	}

	go func() {
		for msg := range msgs {
			var data types.SeedData
			if err := json.Unmarshal(msg.Body, &data); err != nil {
				fmt.Println(err)
				break
			}
			merchantCollection := client.Database("matrix_merchants").Collection("merchant_details")
			merchant := types.Merchant{
				Name:  data.Name,
				Email: data.Email,
			}
			result, err := merchantCollection.InsertOne(context.Background(), merchant)
			if err != nil {
				fmt.Println(err)
				break
			}
			merchantID := result.InsertedID.(primitive.ObjectID)

			for _, pincode := range data.Pincodes {
				databaseName := GetDatabaseName(pincode)
				pincodeCollection := client.Database(databaseName).Collection("pincode_map")
				filter := bson.M{"_id": pincode}
				update := bson.M{
					"$addToSet": bson.M{
						"merchant_ids": merchantID,
					},
				}
				_, err := pincodeCollection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
				if err != nil {
					fmt.Println(err)
					break
				}
			}
		}
	}()

}
