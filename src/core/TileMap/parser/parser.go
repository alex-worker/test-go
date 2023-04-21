package parser

import (
	"encoding/xml"
)

// LoadTmx
// нужен IResourceManager так как могут быть вложенные файлы и их тоже нужно загружать
func LoadTmx(buf *[]byte) (*TmxMap, error) {
	var myMap TmxMap

	err := xml.Unmarshal(*buf, &myMap)
	if err != nil {
		panic(err)
	}

	return &myMap, nil
}
