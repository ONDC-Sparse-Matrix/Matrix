package utils

import (
	"cache-server/types"
)

//TODO: @Dagger will be different for each server 
func InitServer() {
	server := types.ServerInfo{
		Host: "localhost",
		Port: "5000",
	}
	types.Server = server
}
