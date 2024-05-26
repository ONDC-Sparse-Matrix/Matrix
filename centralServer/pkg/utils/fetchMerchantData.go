package utils

import (
	// "../types"
	"bytes"
	"centralCDN/pkg/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func ChooseServerURL(pincode string, clientId string) string {
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
	baseUrl := fmt.Sprintf("http://%s:%s/", cacheServer.Host, cacheServer.Port)
	fmt.Println(baseUrl)
	resp, err := http.Get(baseUrl)

	if err != nil {
		return "Error fetching data"
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return string(body)
}

func FetchMerchantData(pincode string, clientId string) string {
	baseUrl := ChooseServerURL(pincode, clientId)
	baseUrl = baseUrl + "pincode/" + pincode + "/" + clientId
	fmt.Println(baseUrl)
	resp, err := http.Get(baseUrl)

	if err != nil {
		return "Error fetching data"
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return string(body)
}

func UpdateCache(clintId string, cacheResponse []types.PincodeInfo) {
	fmt.Println("Updating cache")
	samplePin := cacheResponse[0].Pincode
	baseUrl := ChooseServerURL(samplePin, clintId)
	jsonData, _ := json.Marshal(cacheResponse)

	resp, err := http.Post(baseUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error in sending the cache response", err)
	}
	defer resp.Body.Close()
}
