package parser

import (
	"encoding/xml"
)

// Parse
func Parse(buf *[]byte) (*TmxMap, error) {
	var myMap TmxMap

	err := xml.Unmarshal(*buf, &myMap)
	if err != nil {
		panic(err)
	}

	return &myMap, nil
}
