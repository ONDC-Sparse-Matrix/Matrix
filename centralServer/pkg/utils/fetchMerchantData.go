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
	// "strconv"
)

func ChooseServerURL(pincode string, clientId string) string {
	var cacheServer types.CacheServer
	// pincodeInt, _ := strconv.Atoi(pincode)
	// regionCode := pincodeInt / 100000
	// switch regionCode {
	// case 1, 2:
		cacheServer = types.CacheServerList[0]
	// case 3, 4:
	// 	cacheServer = types.CacheServerList[1]
	// case 5, 6:
	// 	cacheServer = types.CacheServerList[2]
	// case 7, 8, 9:
	// 	cacheServer = types.CacheServerList[3]
	// default:
	// 	return "Invalid Pincode"
	// }
	baseUrl := fmt.Sprintf("http://%s:%s/", cacheServer.Host, cacheServer.Port)
	fmt.Println(baseUrl)
	// resp, err := http.Get(baseUrl)

	// if err != nil {
	// 	return "36 Error fetching data"
	// }
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	return baseUrl
}

func FetchMerchantData(pincode string, clientId string) string {
	baseUrl := ChooseServerURL(pincode, clientId)
	fmt.Println("hehe = ", baseUrl)
	fmt.Println("ClinentID === ",clientId)
	baseUrl2 := fmt.Sprintf("%spincode/%s/%s",baseUrl,pincode,clientId)
	// baseUrl = baseUrl + "pincode/" + pincode + "/" + clientId
	fmt.Println(baseUrl2)
	resp, err := http.Get(baseUrl2)

	if err != nil {
		fmt.Print("ERROR",err)
		return "51 Error fetching data"
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return string(body)
}

func UpdateCache(clintId string, response types.CachePayload) {
	fmt.Println("Updating cache")
	samplePin := response.CacheResponse[0].Pincode
	baseUrl := ChooseServerURL(samplePin, clintId)
	baseUrl2 := baseUrl + "update/"
	fmt.Println(baseUrl2)
	fmt.Println(response)
	jsonData, _ := json.Marshal(response)
	fmt.Println("Marshaled cache",string(jsonData))

	resp, err := http.Post(baseUrl, "application/", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error in sending the cache response", err)
	}
	defer resp.Body.Close()
	fmt.Println(resp)
}
