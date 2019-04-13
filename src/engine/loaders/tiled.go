package loaders

// https://tutorialedge.net/golang/parsing-xml-with-golang/ 

import (
	"fmt"
	"io/ioutil"
	"encoding/xml"
	"strings"
	"strconv"
	"regexp"
	"../def"
)

func createMap( w uint32, h uint32) [][]def.Cell{

	myMap := make([][]def.Cell, h)

	for i := range myMap {
		myMap[i] = make([]def.Cell, w)
	}

	return myMap

}

func getTileInfo(tsxmap tsxTileSet) (tileName string, w int32, h int32 ) {

	w64, err := strconv.ParseUint(tsxmap.Width, 10, 32)
    if err != nil {
        panic(err)
	}

	w = int32(w64)

	h64, err := strconv.ParseUint(tsxmap.Height, 10, 32)
    if err != nil {
        panic(err)
	}
	
	h = int32(h64)

	tileName = tsxmap.Image.Source

	return 
}

// LoadTSX загружаем файл описания тайлов
func LoadTSX(filename string) (tileName string, w int32, h int32) {
	fmt.Println("Loading TSX...", filename)

	xmlFile, err := def.OpenFile(filename)
	if err != nil {
		panic(err)
	}	
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var tsxmap tsxTileSet
	xml.Unmarshal(byteValue, &tsxmap)

	return getTileInfo(tsxmap)
}

// LoadTmx по файлу возвращаются 
// cells - карта width x height
// tsxFileName - имя файла описания
func LoadTmx(filename string) (cells *[][]def.Cell, tiles *[]def.Tile, tileFileName string, tileW int32, tileH int32 ) {
	fmt.Println("Loading map...", filename)

	xmlFile, err := def.OpenFile(filename)
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

	var tsxFileName = tmxmap.TileSet.Source
	if tsxFileName == "" {
		// тайлы прямо в файле 
		tileFileName, tileW, tileH = getTileInfo( tmxmap.TileSet )
	} else {
		// println( tsxFileName )
		tileFileName, tileW, tileH = LoadTSX( tsxFileName )
	}

	cells = &myMap
	
	return
	
	// return &myMap, w, h, "lol"
	
}