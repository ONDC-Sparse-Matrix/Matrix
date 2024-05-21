package utils

import (
	"cache-server/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	// "net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

const (
	DefaultExpiration = 24 * time.Hour
	PurgeInterval     = 64 * time.Hour
)

// JSON Format to recieve in body from database
// {
// {pincode : 221144, merchantList : ["merchant1", "merchant2", "merchant3"]},
// {pincode : 221144, merchantList : ["merchant1", "merchant2", "merchant3"]},
// {pincode : 221144, merchantList : ["merchant1", "merchant2", "merchant3"]},
// {pincode : 221144, merchantList : ["merchant1", "merchant2", "merchant3"]},
// {pincode : 221144, merchantList : ["merchant1", "merchant2", "merchant3"]},
// {pincode : 221144, merchantList : ["merchant1", "merchant2", "merchant3"]},
// {pincode : 221144, merchantList : ["merchant1", "merchant2", "merchant3"]},
// }

func cacheMiss(pincode string, c *cache.Cache) types.SendPincodeInfo {
	// pincodeInt, _ := strconv.Atoi(pincode)
	// for i := pincodeInt - 5; i <= pincodeInt+5; i++ {
	fmt.Println("Cache miss")
	var pincodeInfo types.PincodeInfo
	var newPincodeInfoList []types.PincodeInfo
	// var currentPincodeInfo types.PincodeInfo
	var sendPincodeInfo3 types.ParseJson
	var sendPincodeInfo types.SendPincodeInfo
	var sendPincodeInfo2 types.SendPincodeInfo

	baseUrl := fmt.Sprintf("http://%s:%s/merchants/%s", types.Server.Host, types.Server.Port, pincode)
	fmt.Println(baseUrl)
	resp, err := http.Get(baseUrl)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	err = json.Unmarshal(body, &(sendPincodeInfo3))
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		// return nil
	}
	sendPincodeInfo = sendPincodeInfo3.SendPincodeInfo2

	currentPincodeInfo := sendPincodeInfo.Current
	var startIndex, endIndex, count int
	startIndex = 20
	endIndex = 40
	fmt.Println("CAche ================== ", sendPincodeInfo.Cache)
	// Access each element in the array
	count = 0
	for _, PincodeInfo := range sendPincodeInfo.Cache {
		var cacheInfo types.CacheInfo

		pincodeInfo.MerchantList = PincodeInfo.MerchantList
		pincodeInfo.Pincode = PincodeInfo.Pincode
		fmt.Println("asdasd ", pincodeInfo, "\n\n")
		count++
		if count >= startIndex && count <= endIndex {
			newPincodeInfoList = append(newPincodeInfoList, pincodeInfo)
		}
		cacheInfo.LastModified = time.Now().UTC()
		cacheInfo.MerchantList = PincodeInfo.MerchantList
		c.Set(pincode, cacheInfo, cache.DefaultExpiration)
		// fmt.Printf("Pincode: %s, Merchants: %v\n", PincodeInfo.Pincode, PincodeInfo.MerchantList)
	}
	fmt.Println(newPincodeInfoList)
	sendPincodeInfo2.Current = currentPincodeInfo
	sendPincodeInfo2.Cache = newPincodeInfoList
	// types.PincodeInfoList = append(types.PincodeInfoList, pincodeInfo)
	// }
	return sendPincodeInfo2
}

func CheckPincode(pincode string, ctx *fiber.Ctx, c *cache.Cache) []byte {
	types.PincodeInfoList = types.PincodeInfoList[:0]
	var sendPincodeInfo types.SendPincodeInfo
	var pincodeInfoList []types.PincodeInfo
	var pincodeInfoo types.PincodeInfo
	if x, found := c.Get(pincode); found {
		fmt.Println("Cache hit")
		var cacheInfo types.CacheInfo
		cacheInfo = x.(types.CacheInfo)

		sendPincodeInfo.Current.Pincode = pincode
		sendPincodeInfo.Current.MerchantList = cacheInfo.MerchantList

		// Fill cache server
		pincodeInt, _ := strconv.Atoi(pincode)
		for i := pincodeInt - 10; i <= pincodeInt+10; i++ {
			if x, found := c.Get(strconv.Itoa(i)); found {
				fmt.Println("Cache hit")
				var cacheInfo types.CacheInfo
				cacheInfo = x.(types.CacheInfo)
				pincodeInfoo.Pincode = pincode
				pincodeInfoo.MerchantList = cacheInfo.MerchantList

				pincodeInfoList = append(pincodeInfoList, pincodeInfoo)
			} else {
				continue
			}
		}
		sendPincodeInfo.Cache = pincodeInfoList
	} else {
		sendPincodeInfo = cacheMiss(pincode, c)
	}
	jsonData, _ := json.Marshal(sendPincodeInfo)
	return jsonData
}

func CheckCache(pincodeList []string, c *cache.Cache) {

	for _, pincode := range pincodeList {
		if x, found := c.Get(pincode); found {
			fmt.Println("Cache hit now update")
			cacheInfo := x.(types.CacheInfo)
			cacheInfo.MerchantList = append(cacheInfo.MerchantList)
		}
	}
}
