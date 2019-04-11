package loaders

// https://tutorialedge.net/golang/parsing-xml-with-golang/ 

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/xml"
	"strings"
	"strconv"
	"regexp"
	"../def"
)

type tmxMap struct {
	XMLName xml.Name `xml:"map"`
	Width string `xml:"width,attr"`
	Height string `xml:"height,attr"`
	TileWidth string `xml:"tilewidth,attr"`
	TileHeight string `xml:"tileheight,attr"`
    Layer      tmxLayer `xml:"layer"`
}

type tmxLayer struct {
	XMLName xml.Name `xml:"layer"`
	Name string `xml:"name,attr"`
	Width string `xml:"width,attr"`
	Height string `xml:"height,attr"`
	Data string `xml:"data"`
}

func createMap( w uint32, h uint32) [][]def.Cell{

	myMap := make([][]def.Cell, h)

	for i := range myMap {
		myMap[i] = make([]def.Cell, w)
	}

	return myMap

}

// LoadTmxMap по файлу возвращаются 
// cells - карта width x height
// tsxFileName - имя файла описания
func LoadTmx(filename string) (cells *[][]def.Cell, tsxFileName string ) {
	fmt.Println("Loading map...")

	xmlFile, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var tmxmap tmxMap

	xml.Unmarshal(byteValue, &tmxmap)

	re := regexp.MustCompile(`\r?\n`)
	normalizedMap := re.ReplaceAllString(tmxmap.Layer.Data, "")

	myMapStr := strings.Split(normalizedMap, ",")

	w64, err := strconv.ParseUint(tmxmap.Width, 10, 32)
    if err != nil {
        panic(err)
	}

	// var w uint32
	// var h uint32

	w := uint32(w64)

	h64, err := strconv.ParseUint(tmxmap.Height, 10, 32)
    if err != nil {
        panic(err)
	}
	
	h := uint32(h64)

	myMap := createMap( w, h )

	var x uint32
	var y uint32
	for _, c := range myMapStr {
		cell, err := strconv.ParseUint(c, 10, 32)
		if err != nil {
			panic(err)
		}
		myMap[y][x] = def.Cell(cell-1)
		x++
		if ( x == w ){
			y++
			x = 0
		}
	}

	cells = &myMap
	
	return
	
	// return &myMap, w, h, "lol"
	
}