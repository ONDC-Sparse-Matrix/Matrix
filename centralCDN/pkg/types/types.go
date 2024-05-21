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
