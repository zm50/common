package serialize

import "gopkg.in/yaml.v3"

type YamlSerializer struct {}

func (y *YamlSerializer) Serialize(data interface{}) ([]byte, error) {
	return yaml.Marshal(data)
}

func (y *YamlSerializer) Deserialize(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}
