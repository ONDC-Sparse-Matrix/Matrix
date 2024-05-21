package utils

import (
	"cache-server/types"
)

func InitServer() {
	server := types.ServerInfo{
		Host: "localhost",
		Port: "5000",
	}
	types.Server = server
}
