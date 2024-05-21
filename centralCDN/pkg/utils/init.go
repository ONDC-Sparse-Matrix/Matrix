package utils

import (
	"fmt"

	"centralCDN/pkg/types"
)

func InitCacheServerList() {
	// var cacheServerList types.CacheServerList
	types.CacheServerList = append(types.CacheServerList, types.CacheServer{Host: "localhost", Port: "4000"})
	types.CacheServerList = append(types.CacheServerList, types.CacheServer{Host: "localhost", Port: "8001"})
	types.CacheServerList = append(types.CacheServerList, types.CacheServer{Host: "localhost", Port: "8002"})
	types.CacheServerList = append(types.CacheServerList, types.CacheServer{Host: "localhost", Port: "8003"})
}

func InitPincode() {
	types.PincodeCount = 30000
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
	types.ServerRangeList = append(types.ServerRangeList, types.ServerRange{RangeStart: 1, RangeEnd: 2, CacheServerID: types.CacheServerList[0]})
	types.ServerRangeList = append(types.ServerRangeList, types.ServerRange{RangeStart: 3, RangeEnd: 4, CacheServerID: types.CacheServerList[1]})
	types.ServerRangeList = append(types.ServerRangeList, types.ServerRange{RangeStart: 5, RangeEnd: 6, CacheServerID: types.CacheServerList[2]})
	types.ServerRangeList = append(types.ServerRangeList, types.ServerRange{RangeStart: 7, RangeEnd: 9, CacheServerID: types.CacheServerList[3]})

}
