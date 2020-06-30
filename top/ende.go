package top

import "encoding/json"

//FormatData ...
func FormatData(top interface{}) ([]byte, error) {
	jsonData, err := json.MarshalIndent(top, "", "  ")
	if err != nil {
		return jsonData, err
	}
	return jsonData, nil
}
