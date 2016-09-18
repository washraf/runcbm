package common

import (
	"encoding/json"
	"errors"
	"fmt"
	//"fmt"
	"math"
	"os/exec"
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
			it := items[i]
			rss := data[it].([]interface{})
			//fmt.Println(rss[0])
			s, ok = rss[0].(map[string]interface{})
			if !ok {
				return "", errors.New("item is not a map")
			}
		}
		data = s
	}
	last := fmt.Sprint(data[items[len(items)-1]])
	//last, ok := data[items[len(items)-1]].(string)
	//fmt.Println(last)
	//if !ok {
	//	return "", errors.New("item is not a string")
	//}
	return last, nil
}

//ReadUsingCrit ...
func ReadUsingCrit(condir string, fileName string, middle string, minors ...string) ([]string, error) {
	command := exec.Command("crit", "show", fileName)
	command.Dir = condir
	res, err := command.CombinedOutput()

	if err != nil {
		//fmt.Println("error line 59")
		return nil, err
	}
	var strArr []string
	for _, i := range minors {
		r, err := GetItemFromJSON(res, "entries", middle, i)
		if err != nil {
			//fmt.Println("error line 66")
			return nil, err
		}
		strArr = append(strArr, r)
		//fmt.Println("Walid:", i, r)
	}
	return strArr, nil
}
