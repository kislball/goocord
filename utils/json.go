package utils

import "encoding/json"

// Parse is a wrapper around json.Unmarshal
func Parse(data string) (st interface{}, err error) {
	err = json.Unmarshal([]byte(data), &st)
	return
}

// Stringify is a wrapper around json.Marshal
func Stringify(data interface{}) (d string, err error) {
	b, err := json.Marshal(data)
	d = string(b)
	return
}
