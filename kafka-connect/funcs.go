package main

import (
	"encoding/json"
	"time"
)

func jsonToMap(byteValue []byte) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	if err := json.Unmarshal(byteValue, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func unixTimeToTime(unixTime int64) time.Time {
	return time.UnixMicro(unixTime).UTC()
}
