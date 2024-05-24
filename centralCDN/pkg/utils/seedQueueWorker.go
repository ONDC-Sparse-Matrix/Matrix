package utils

import (
	"centralCDN/pkg/types"
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HandleSeedQueue(queueConnection *amqp.Connection,client *mongo.Client){
	msgs,err := ConsumeSeeding(queueConnection)
	if err != nil {
		fmt.Println(err)
	}

	for msg := range msgs {
		var data types.SeedData
		if err := json.Unmarshal(msg.Body, &data); err != nil {
			break
		}
		merchantCollection := client.Database("matrix_merchants").Collection("merchant_details")
		merchant := types.Merchant{
			Name:  data.Name,
			Email: data.Email,
		}
		result, err := merchantCollection.InsertOne(context.Background(), merchant)
		if err != nil {
			break
		}
		merchantID := result.InsertedID.(primitive.ObjectID)

		for _, pincode := range data.Pincodes {
			databaseName := GetDatabaseName(pincode)
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
}