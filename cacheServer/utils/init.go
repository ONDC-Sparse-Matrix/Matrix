package utils

import (
	"cache-server/types"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//TODO: @Dagger will be different for each server

func InitServer() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("INSTANCE_HOST_IP")

	server := types.ServerInfo{
		Host: host,
		Port: "5000",
	}
	types.Server = server
}
