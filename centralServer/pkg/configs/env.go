package configs

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

func EnvMongoURI() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv("MONGOURI")
}

func EnvFrontendURL() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv("FRONTEND_URL")
}

func EnvRabbitMQURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("RABBITMQ_URI")
}