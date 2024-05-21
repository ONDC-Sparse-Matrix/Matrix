package utils

import (
	// "../types"
	"centralCDN/pkg/types"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func FetchMerchantData(pincode string) string {
	var cacheServer types.CacheServer
	pincodeInt, _ := strconv.Atoi(pincode)
	regionCode := pincodeInt / 100000
	switch regionCode {
	case 1, 2:
		cacheServer = types.CacheServerList[0]
	case 3, 4:
		cacheServer = types.CacheServerList[1]
	case 5, 6:
		cacheServer = types.CacheServerList[2]
	case 7, 8, 9:
		cacheServer = types.CacheServerList[3]
	default:
		return "Invalid Pincode"
	}
	baseUrl := fmt.Sprintf("http://%s:%s/pincode/%s", cacheServer.Host, cacheServer.Port, pincode)
	resp, err := http.Get(baseUrl)
	fmt.Println(baseUrl)
	if err != nil {
		return "Error fetching data"
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return string(body)
}
