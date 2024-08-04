package serialize

import "encoding/xml"

type XmlSerializer struct {}

func (s *XmlSerializer) Marshal(data interface{}) ([]byte, error) {
	return xml.Marshal(data)
}

func (s *XmlSerializer) Unmarshal(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}
