package world

// https://tutorialedge.net/golang/parsing-xml-with-golang/ 

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/xml"
	"strings"
	"strconv"
	"regexp"
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

func createMap( w uint32, h uint32) [][]Cell{

	myMap := make([][]Cell, h)

	for i := range myMap {
		myMap[i] = make([]Cell, w)
	}

	return myMap

}

func loadTmxMap(filename string) (*[][]Cell, uint32, uint32) {
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
	var w uint32 = uint32(w64)

	h64, err := strconv.ParseUint(tmxmap.Height, 10, 32)
    if err != nil {
        panic(err)
	}
	var h uint32 = uint32(h64)

	myMap := createMap( w, h )

	var x uint32 = 0
	var y uint32 = 0
	for _, c := range myMapStr {
		cell, err := strconv.ParseUint(c, 10, 32)
		if err != nil {
			panic(err)
		}
		myMap[y][x] = Cell(cell-1)
		x++
		if ( x == w ){
			y++
			x = 0
		}
	}

	return &myMap, w, h
	
}