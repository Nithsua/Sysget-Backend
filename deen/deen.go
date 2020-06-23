package deen

import (
	"encoding/json"
)

//FormatData ...
func FormatData(data interface{}) ([]byte, error) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return jsonData, err
	}
	return jsonData, nil
}
