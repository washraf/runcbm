package common

import (
	"encoding/json"
	"errors"
	"math"
)

//Exit ...
type Exit struct {
	Code int
}

//Round ...
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

//GetItemFromJSON ...
func GetItemFromJSON(rawdata []byte, items ...string) (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal(rawdata, &data)
	if err != nil {
		return "", err
	}
	for i := 0; i < len(items)-1; i++ {
		s, ok := data[items[i]].(map[string]interface{})
		if !ok {
			return "", errors.New("item is not a map")
		}
		data = s
	}
	last, ok := data[items[len(items)-1]].(string)
	if !ok {
		return "", errors.New("item is not a string")
	}
	return last, nil
}
