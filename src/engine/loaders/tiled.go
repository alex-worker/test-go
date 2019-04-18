package loaders

// https://tutorialedge.net/golang/parsing-xml-with-golang/

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"../def"
)

// TileSetInfo описание тайлсета для дальнейшей обработки
type TileSetInfo struct {
	Filename string
	Tiles *def.AnimateTiles
	TileW    int32
	TileH    int32
}

func createMap(w uint32, h uint32) [][]def.Cell {

	myMap := make([][]def.Cell, h)

	for i := range myMap {
		myMap[i] = make([]def.Cell, w)
	}

	return myMap

}

func parseLayer(layer *tmxLayer) *def.Layer {

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

	myMap := createMap(w, h)

	var x uint32
	var y uint32
	for _, c := range myMapStr {
		cell, err := strconv.ParseUint(c, 10, 32)
		if err != nil {
			panic(err)
		}
		myMap[y][x] = def.Cell(cell - 1)
		x++
		if x == w {
			y++
			x = 0
		}
	}

	return &def.Layer{
		Data: &myMap,
		Name: layer.Name,
		W:    w,
		H:    h,
	}

}

func parseTileSet(tileset *tsxTileSet) TileSetInfo {

	var tsxFileName = tileset.Source
	if tsxFileName != "" {
		return loadTSX(tsxFileName)
	}

	w64, err := strconv.ParseUint(tileset.TileWidth, 10, 32)
	if err != nil {
		panic(err)
	}

	w := int32(w64)

	h64, err := strconv.ParseUint(tileset.TileHeight, 10, 32)
	if err != nil {
		panic(err)
	}

	h := int32(h64)

	tileName := tileset.Image.Source

	myTileSet := TileSetInfo{
		Filename: tileName,
		TileW:    w,
		TileH:    h,
	}

	fmt.Println("TILESET", tileset.Name, "parsed OK")

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
// layers - карта width x height
// tsxFileName - имя файла описания
func LoadTmx(filename string) (mymap *def.Map, tsetsPtr *[]TileSetInfo) {
	fmt.Println("Loading map...", filename)

	xmlFile, err := def.OpenFile(filename)
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var tmxmap tmxMap

	xml.Unmarshal(byteValue, &tmxmap)

	lenLayers := len(tmxmap.Layers)
	layers := def.Layers(make([]def.Layer, lenLayers))

	for i, layer := range tmxmap.Layers {
		curlayer := parseLayer(layer)
		layers[i] = *curlayer
	}

	tilesets := make([]TileSetInfo, len(tmxmap.TileSets) )

	setlen := len(tmxmap.TileSets)

	for i, tileset := range tmxmap.TileSets {
		name := fmt.Sprint(setlen - i)
		fmt.Println(tileset.Name)
		tilesets[i] = parseTileSet(tileset)
		fmt.Println( "tileset name:", name )
	}

	mymap = &def.Map{
		Layers: layers,
		W:      layers[0].W, // ориентируемся по первому слою ( земля )
		H:      layers[0].H,
	}
	tsetsPtr = &tilesets

	return

}
