package parser

import (
	"encoding/xml"
	"fmt"
	. "test-go/src/core/FileManager"
	. "test-go/src/interfaces/IResourceManager"
)

// LoadTmx
// нужен IResourceManager так как могут быть вложенные файлы и их тоже нужно загружать
func LoadTmx(r *IResourceManager, filename string) (*TmxMap, error) {
	fmt.Println("Loading map...", filename)

	buf, err := GetFile(r, filename)
	if err != nil {
		return nil, err
	}

	var myMap TmxMap

	err = xml.Unmarshal(*buf, &myMap)
	if err != nil {
		panic(err)
	}

	return &myMap, nil
}
