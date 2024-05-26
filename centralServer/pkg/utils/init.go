package utils

import (
	"fmt"
	"log"
	"os"

	"centralCDN/pkg/types"

	"github.com/joho/godotenv"
)

func InitCacheServerList() {
	// var cacheServerList types.CacheServerList
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host1 := os.Getenv("CACHE_SERVER_HOST_1")
	// host2 := os.Getenv("CACHE_SERVER_HOST_2")
	// host3 := os.Getenv("CACHE_SERVER_HOST_3")
	// host4 := os.Getenv("CACHE_SERVER_HOST_4")

	port1 := os.Getenv("CACHE_SERVER_PORT_1")
	// port2 := os.Getenv("CACHE_SERVER_PORT_2")
	// port3 := os.Getenv("CACHE_SERVER_PORT_3")
	// port4 := os.Getenv("CACHE_SERVER_PORT_4")

	types.CacheServerList = append(types.CacheServerList, types.CacheServer{Host: host1, Port: port1})
	// types.CacheServerList = append(types.CacheServerList, types.CacheServer{Host: host2, Port: port2})
	// types.CacheServerList = append(types.CacheServerList, types.CacheServer{Host: host3, Port: port3})
	// types.CacheServerList = append(types.CacheServerList, types.CacheServer{Host: host4, Port: port4})
}

func InitPincode() {
	types.PincodeCount = 30000
	types.MinFreq = 0
}

func InitServerRangeList() {
	numberOfCacheServers := len(types.CacheServerList)
	fmt.Println(numberOfCacheServers)

	// for i := 0; i < numberOfCacheServers-1; i++ {
	// rangeStart := (types.PincodeCount / numberOfCacheServers) * i
	// rangeEnd := (types.PincodeCount / numberOfCacheServers) * (i + 1)

	// fmt.Println(rangeStart, rangeEnd)

	// serverRange :=
	// types.ServerRangeList = append(types.ServerRangeList, types.ServerRange{RangeStart: , RangeEnd: , CacheServerID: types.CacheServerList[i]})
	// }
	types.ServerRangeList = append(types.ServerRangeList, types.ServerRange{RangeStart: 1, RangeEnd: 9, CacheServerID: types.CacheServerList[0]})
	// types.ServerRangeList = append(types.ServerRangeList, types.ServerRange{RangeStart: 3, RangeEnd: 4, CacheServerID: types.CacheServerList[1]})
	// types.ServerRangeList = append(types.ServerRangeList, types.ServerRange{RangeStart: 5, RangeEnd: 6, CacheServerID: types.CacheServerList[2]})
	// types.ServerRangeList = append(types.ServerRangeList, types.ServerRange{RangeStart: 7, RangeEnd: 9, CacheServerID: types.CacheServerList[3]})

}
