package types

import "time"

type ServerInfo struct {
	Host string
	Port string
}

type CacheInfo struct {
	LastModified time.Time
	MerchantList []MerchantInfo
}
type MerchantInfo struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	PinCodes []string `json:"pin_codes"`
}

type PincodeInfo struct {
	Pincode      string         `json:"pincode"`
	MerchantList []MerchantInfo `json:"merchantList"`
}
type SendPincodeInfo struct {
	Current PincodeInfo   `json:"current"`
	Cache   []PincodeInfo `json:"cache"`
}
type ParseJson struct {
	Status           int             `json:"status"`
	Message          string          `json:"message"`
	SendPincodeInfo2 SendPincodeInfo `json:"data"`
}

var PincodeInfoList []PincodeInfo

// var data struct {
// 	InfoList []PincodeInfo `json:"PincodeInfo"`
// }

var Server ServerInfo
