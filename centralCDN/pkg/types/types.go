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
