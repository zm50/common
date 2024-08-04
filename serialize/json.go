package serialize

import "encoding/json"

type JsonSerializer struct {}

func (j JsonSerializer) Marshal(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func (j JsonSerializer) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
