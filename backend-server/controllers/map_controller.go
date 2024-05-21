package controllers

import (
	// "context"
	// "net/http"
	// "regional_server/configs"
	// "regional_server/models"
	// "regional_server/responses"
	// "time"

	"context"
	"net/http"
	"regional_server/models"
	"regional_server/responses"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "fmt"
)

func UpdateMap(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// pinCode := c.Params("pincode")

	var merchants models.Map
	err := mapCollection.FindOne(ctx, bson.M{}).Decode(&merchants)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.Response{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}
	return c.Status(http.StatusOK).JSON(responses.Response{
		Data: &fiber.Map{"data": merchants},
	})
}
