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

func loadMap(filename string) (*[]Cell, uint32, uint32) {
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
	
	w, err := strconv.ParseUint(tmxmap.Width, 10, 32)
    if err != nil {
        panic(err)
	}
	
	h, err := strconv.ParseUint(tmxmap.Height, 10, 32)
    if err != nil {
        panic(err)
	}

	myMap := make([]Cell, w * h )

	for i, c := range myMapStr {
		cell, err := strconv.ParseUint(c, 10, 32)
		// fmt.Println( err, i, cell )
		if err != nil {
			panic(err)
		}	
		myMap[i] = Cell( cell ) 
	}

	return &myMap, uint32(w), uint32(h)
	
}