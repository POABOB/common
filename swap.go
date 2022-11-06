package utils

import (
	"encoding/json"
)

// 通過 json tag 對 struct 賦值
func SwapTo(request, target interface{}) (err error) {
	dataByte, err := json.Marshal(request)
	if err != nil {
		return
	}
	err = json.Unmarshal(dataByte, target)
	return
}
