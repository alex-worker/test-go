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

// TileSetInfo описание тайлсета
type TileSetInfo struct {
	Filename string
	TileW int32
	TileH int32
}

func createMap( w uint32, h uint32) [][]def.Cell{

	myMap := make([][]def.Cell, h)

	for i := range myMap {
		myMap[i] = make([]def.Cell, w)
	}

	return myMap

}

func parseLayer( layer *tmxLayer ) *[][]def.Cell {

	re := regexp.MustCompile(`\r?\n`)
	normalizedMap := re.ReplaceAllString(layer.Data, "")
	myMapStr := strings.Split(normalizedMap, ",")

	w64, err := strconv.ParseUint(layer.Width, 10, 32)
    if err != nil {
        panic(err)
	}

	w := uint32(w64)

	h64, err := strconv.ParseUint(layer.Height, 10, 32)
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
	
	return &myMap

}

func parseTileSet(tileset *tsxTileSet) TileSetInfo {

	var tsxFileName = tileset.Source
	if ( tsxFileName!= ""){
		return loadTSX(tsxFileName)
	}

	w64, err := strconv.ParseUint(tileset.Width, 10, 32)
    if err != nil {
        panic(err)
	}

	w := int32(w64)

	h64, err := strconv.ParseUint(tileset.Height, 10, 32)
    if err != nil {
        panic(err)
	}
	
	h := int32(h64)

	tileName := tileset.Image.Source

	myTileSet := TileSetInfo{ 
		Filename: tileName,
		TileW: w,
		TileH: h,
	}

	fmt.Println( "TILESET", tileset.Name, "parsed OK")
	
	return myTileSet
}

func loadTSX(filename string) TileSetInfo {
	fmt.Println("Loading TSX...", filename)

	xmlFile, err := def.OpenFile(filename)
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var tsxmap tsxTileSet
	xml.Unmarshal(byteValue, &tsxmap)

	return parseTileSet(&tsxmap)
}

// LoadTmx по файлу возвращаются 
// cells - карта width x height
// tsxFileName - имя файла описания
func LoadTmx(filename string) (cells *[][]def.Cell, tsets *map[string]TileSetInfo ) {
	fmt.Println("Loading map...", filename)

	xmlFile, err := def.OpenFile(filename)
	if err != nil {
		panic(err)
	}	
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var tmxmap tmxMap
	var layers *[][]def.Cell
	// var tilesets []*TileSetInfo

	xml.Unmarshal(byteValue, &tmxmap)

	for _, layer := range tmxmap.Layers {
		layers = parseLayer( layer )
	}

	// tilesetLen := len( tmxmap.Layers )
	
	tilesets := make(map[string]TileSetInfo)

	for i, tileset := range tmxmap.TileSets {
		fmt.Println(i)
		tilesets[tileset.Name] = parseTileSet( tileset )
	}

	cells = layers
	tsets = &tilesets

	return
	
}