package utils

import (
	"centralCDN/pkg/types"
	"fmt"
)

func UpdateFreqMap(pincode int) {

	fmt.Println("Updating frequency map with pincode ", pincode)

	freq := types.FrequencyMap[pincode]
	types.FrequencyMap[pincode] = types.FrequencyMap[pincode] + 1

	if len(types.Top50) < 50 {
		fmt.Println("Less than 50 in map ", pincode)
		types.Top50[freq+1] = append(types.Top50[freq+1], pincode)
		if freq+1 > types.MinFreq {
			types.MinFreq = freq + 1
		}
	} else {
		fmt.Println("More than 50 ", pincode)
		if freq+1 < types.MinFreq {
			types.Top50[freq+1] = append(types.Top50[freq+1], pincode)
			delete(types.Top50, types.MinFreq)
			types.MinFreq = freq + 1
			fmt.Println("Updating top 50 ", pincode)
		}
	}

	for key, array := range types.Top50 {
        fmt.Printf("Array for key %d: %v\n", key, array)
    }

}