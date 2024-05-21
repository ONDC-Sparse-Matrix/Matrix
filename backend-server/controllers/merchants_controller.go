package controllers

import (
	"bytes"
	"context"
	"encoding/json"
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

var mapCollection *mongo.Collection = configs.GetCollection(configs.DB, "maps")
var merchantsCollection *mongo.Collection = configs.GetCollection(configs.DB, "merchants")

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
	MerchantList []NewMerchant `json:"merchantList"`
}

var PincodeInfoList []PincodeInfo

func GetMerchants(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	pinCode := c.Params("pincode")
	fmt.Println(pinCode)

	num, err := strconv.ParseFloat(pinCode, 64)
	//??
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}
	cacheResponse := make([]PincodeInfo, 0)
	// finding cache responses
	fmt.Println(num)
	for i := num - 30; i < num+30; i++ {
		fmt.Println("heheheh = ", i)
		var cacheMerchants models.Map
		formatedString := strconv.FormatFloat(i, 'f', -1, 64)
		println(formatedString)
		err = mapCollection.FindOne(ctx, bson.M{"pin_code": formatedString}).Decode(&cacheMerchants)
		if err != nil {
			fmt.Println(cacheMerchants, " Not Found")
			fmt.Println(err)
			continue
			// return c.Status(http.StatusInternalServerError).JSON(responses.Response{
			// 	Status:  http.StatusInternalServerError,
			// 	Message: "error",
			// 	Data:    &fiber.Map{"data": err.Error()},
			// })
		}
		cacheArr := cacheMerchants.MERCHANT_IDS
		fmt.Println(cacheArr)
		cacheSingleresponse := make([]NewMerchant, 0)
		for _, cacheR := range cacheArr {
			var cacheM NewMerchant
			objID, err := primitive.ObjectIDFromHex(cacheR)
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(responses.Response{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    &fiber.Map{"data": err.Error()},
				})
			}
			// fmt.Println(cacheR)
			err = merchantsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&cacheM)
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(responses.Response{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    &fiber.Map{"data": err.Error()},
				})
			}
			cacheSingleresponse = append(cacheSingleresponse, cacheM)
		}

		iStrstr := strconv.FormatFloat(i, 'f', -1, 64)
		var pincodeInfo PincodeInfo
		pincodeInfo.MerchantList = cacheSingleresponse
		pincodeInfo.Pincode = iStrstr
		fmt.Println("pincode = ", pincodeInfo)
		cacheResponse = append(cacheResponse, pincodeInfo)
	}
	fmt.Println("CACHE RESPONSE", cacheResponse)
	// finding current response
	var merchants models.Map
	err = mapCollection.FindOne(ctx, bson.M{"pin_code": pinCode}).Decode(&merchants)
	fmt.Println("ERROR", err)
	fmt.Println("CURRENT RESPONSE", merchants)
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
			"cache": cacheResponse,
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
				bson.M{"pin_code": pinCode},
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
