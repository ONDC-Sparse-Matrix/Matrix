package types

type CacheServer struct {
	Host string
	Port string
}

var CacheServerList []CacheServer

var PincodeCount int

type ServerRange struct {
	CacheServerID CacheServer
	RangeStart    int
	RangeEnd      int
}

var ServerRangeList []ServerRange

// array of pincode vs frequency
var FrequencyMap = make(map[int]int)

var MinFreq int

// Array of frequency vs pincode
var Top50 = make(map[int][]int)

type Merchant struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PincodeInfo struct {
	Pincode      string        `json:"pincode"`
	MerchantList []Merchant `json:"merchantList"`
}

type CachePayload struct{
	CacheResponse []PincodeInfo `json:"cacheResponse"`
	ClientCacheResponse []PincodeInfo `json:"clientCacheResponse"`
}

type SeedData struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Pincodes []int32 `json:"pincodes"`
}