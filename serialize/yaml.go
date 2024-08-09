package serialize

import "gopkg.in/yaml.v3"

type YamlSerializer struct {}

func (y *YamlSerializer) Marshal(data interface{}) ([]byte, error) {
	return yaml.Marshal(data)
}

func (y *YamlSerializer) Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}
