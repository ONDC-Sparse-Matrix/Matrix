package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"regional_server/configs"
	"regional_server/models"
	"regional_server/responses"

	// "net/url"
	"strconv"
	"time"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var mapCollection *mongo.Collection = configs.DB.Database("matrix_map_1").Collection("pincode_map")
var merchantsCollection *mongo.Collection = configs.DB.Database("matrix_merchants").Collection("merchant_details") //TODO: @Wayne write logic to get collection names

type NewMerchant struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	PinCodes []string `json:"pin_codes"`
}

type UpdateMerchant struct {
	ObjectId string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
type PincodeInfo struct {
	Pincode      string        `json:"pincode"`
	MerchantList []models.Merchant `json:"merchantList"`
}

var PincodeInfoList []PincodeInfo

func Send_SSE_Caching_responses (num int,clientId string) {
	ctx,cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	//TODO: @Garv currently the cache_range is hardcoded , but we have to apply some algorithm to calculate the cache_range
	//To Central server
	central_cache_range := 5
	cacheResponse := make([]PincodeInfo, 0)
	//! hande the case when num is less than central_cache_range
	for i :=num-central_cache_range; i<num+central_cache_range; i++{
		log.Println("i",i)
		var current_pincode_map models.Map
		formattedString := strconv.FormatInt(int64(i),10)
		err := mapCollection.FindOne(context.Background(), bson.M{"_id": i}).Decode(&current_pincode_map)
		if err != nil {
			fmt.Println(i, "th Pincode data Not Found")
			fmt.Println(err)
			continue
		}
		merchant_ids_arr := current_pincode_map.MERCHANT_IDS
		//get the details of the merchants from their ids
		merchant_details := make([]models.Merchant, 0)
		for _, merchant_id := range merchant_ids_arr{
			var merchant models.Merchant
			objID, err := primitive.ObjectIDFromHex(merchant_id)
			if err != nil {
				log.Println("Error in converting to ObjectID")
			}
			err = merchantsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&merchant)
			if err != nil {
				log.Println("Error in fetching merchant details")
			}
			merchant_details = append(merchant_details, merchant)
		}
		pincodeInfo := PincodeInfo{
			Pincode:      formattedString,
			MerchantList: merchant_details,
		}
		cacheResponse = append(cacheResponse, pincodeInfo)
	}
	log.Println("Cache Response",cacheResponse)

	requestData := map[string]interface{}{
		"cacheResponse": cacheResponse,
		"clientCacheResponse": cacheResponse,
	}
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		log.Println("Error in marshalling the data")
	}
	fmt.Println("request data",requestData)
	resp, err := http.Post(configs.EnvCentralServerURI()+"/cache/"+clientId, "application/json", bytes.NewBuffer(jsonData)) 
	if err != nil {
		log.Println("Error in sending the cache response",err)
	}
	defer resp.Body.Close()
	fmt.Println("Cache Response sent successfully.")
}

func GetMerchants(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)//! reduce the context time here , as Send_SSE_Caching_responses is taking more time
	defer cancel()

	pinCode := c.Params("pincode")
	clientID := c.Params("clientID") //TODO: @Garv generate an uniqueID for every client and send this with the request
	fmt.Println("pincode",pinCode)
	fmt.Println("ClientID",clientID)

	num, err := strconv.ParseInt(pinCode, 10, 64)
	
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "Incorrect Pincode", //TODO: @Garv for this error show a message in the frontend
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	go Send_SSE_Caching_responses(int(num),clientID)

	// finding current response
	var merchants models.Map
	err = mapCollection.FindOne(ctx, bson.M{"_id": num}).Decode(&merchants)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	arr := merchants.MERCHANT_IDS
	length := len(arr)
	fmt.Println(length)

	response := make([]models.Merchant, 0)
	for i := 0; i < length; i++ {
		var merchant models.Merchant
		objID, err := primitive.ObjectIDFromHex(arr[i])
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
		}

		err = merchantsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&merchant)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
		}

		response = append(response, merchant)
	}

	return c.Status(http.StatusOK).JSON(responses.Response{
		Data: &fiber.Map{
			"current": &fiber.Map{
				"pincode":      pinCode,
				"merchantList": response,
			},
			// "cache": cacheResponse, 
		},
	})

	
}

func AddMerchants(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var newMerchants []NewMerchant

	if err := c.BodyParser(&newMerchants); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	for _, newMerchant := range newMerchants {
		merchant := models.Merchant{
			Name:  newMerchant.Name,
			Email: newMerchant.Email,
		}

		insertResult, err := merchantsCollection.InsertOne(ctx, merchant)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
		}

		for _, pinCode := range newMerchant.PinCodes {
			objID, _ := primitive.ObjectIDFromHex(insertResult.InsertedID.(primitive.ObjectID).Hex())
			_, err := mapCollection.UpdateOne(
				ctx,
				bson.M{"_id": pinCode},
				bson.M{"$push": bson.M{"merchant_ids": objID}},
			)

			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(responses.Response{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    &fiber.Map{"data": err.Error()},
				})
			} else {
				apiUrl := "http://" //!
				requestData := map[string]interface{}{
					"pinCodes": newMerchant.PinCodes,
				}

				jsonData, err := json.Marshal(requestData)
				if err != nil {
					return c.Status(http.StatusInternalServerError).JSON(responses.Response{
						Status:  http.StatusInternalServerError,
						Message: "error",
						Data:    &fiber.Map{"data": err.Error()},
					})
				}

				resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(jsonData))
				if err != nil {
					return c.Status(http.StatusInternalServerError).JSON(responses.Response{
						Status:  http.StatusInternalServerError,
						Message: "error",
						Data:    &fiber.Map{"data": err.Error()},
					})
				}
				defer resp.Body.Close()

				fmt.Println("PinCodes sent successfully.")

			}
		}
	}

	return c.Status(http.StatusOK).JSON(responses.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": "Merchants added successfully"},
	})
}


func Test(c *fiber.Ctx) error {
	collection := configs.DB.Database("matrix_map_1").Collection("pincode_map")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var data any
	err := collection.FindOne(ctx, bson.M{"_id":"151505"}).Decode(&data)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}
	return c.Status(http.StatusOK).JSON(responses.Response{
		Data: &fiber.Map{"data": data},
	})

}
// func UpdateMerchant(c *fiber.Ctx)error{
// 	/*
// 	request ->
// 	[
// 		{objectId:, name:,email:}
// 		....
// 	]
// 	itterate -> replace
// 	*/
// 	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	// defer cancel()

// 	// var updatesArray [] UpdateMerchant

// 	// if err := c.BodyParser(&updatesArray);err !=nil{
// 	// 	return c.Status(http.StatusBadRequest).JSON(responses.Response{
// 	// 		Status:  http.StatusBadRequest,
// 	// 		Message: "error",
// 	// 		Data:    &fiber.Map{"data": err.Error()},
// 	// 	})
// 	// }

// 	// for _,update := range updatesArray {
// 	// 	updatedMerchant := models.Merchant{
// 	// 		Name: update.Name,
// 	// 		Email: update.Email,
// 	// 	}

// 	// }

// 	return nil

// }
