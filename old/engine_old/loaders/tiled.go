package loaders

// https://tutorialedge.net/golang/parsing-xml-with-golang/

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"test-go/src/engine/defines"
	"test-go/src/engine/resource"
)

// TileSetInfo описание тайлсета для дальнейшей обработки
type TileSetInfo struct {
	Filename string
	Tiles    *defines.AnimateTiles
	TileW    int32
	TileH    int32
}

func createMap(w uint32, h uint32) [][]defines.Cell {

	myMap := make([][]defines.Cell, h)

	for i := range myMap {
		myMap[i] = make([]defines.Cell, w)
	}

	return myMap

}

func parseLayer(layer *tmxLayer) *defines.Layer {

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
		if cell == 0 {
			myMap[y][x] = 0
		} else {
			myMap[y][x] = defines.Cell(cell - 1)
		}
		x++
		if x == w {
			y++
			x = 0
		}
	}

	return &defines.Layer{
		Data: &myMap,
		Name: layer.Name,
		W:    w,
		H:    h,
	}

}

func loadFrames(t *tsxTile) *[]defines.AnimateFrame {

	f := t.Animations.Frames
	frames := make([]defines.AnimateFrame, len(f))
	for i, tsxFrame := range f {
		frames[i] = defines.AnimateFrame{
			Cell:     defines.Cell(tsxFrame.Tileid),
			Duration: tsxFrame.Duration,
		}
		// fmt.Println( frames[i] )
	}

	return &frames
}

// парсим xml-тайлсет
func parseAnimateTiles(tileset *tsxTileSet) *defines.AnimateTiles {
	fmt.Println("parse animate tileset")

	tiles := make(defines.AnimateTiles)

	for _, tile := range tileset.Tiles {

		animCell := defines.Cell(tile.ID)

		tiles[animCell] = &defines.AnimateTile{
			Tick:   0,
			Index:  0,
			Frames: *loadFrames(tile),
		}
		// fmt.Println("anim cell:", animCell)
	}
	return &tiles
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
		Tiles:    parseAnimateTiles(tileset),
		TileW:    w,
		TileH:    h,
	}

	fmt.Println("TILESET", tileset.Name, "parsed OK")

	return myTileSet
}

func loadTSX(filename string) TileSetInfo {
	fmt.Println("Loading TSX...", filename)

	xmlFile, err := resource.OpenFile(filename)
	if err != nil {
		panic(err)
	}

	defer resource.CloseFile(xmlFile)

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var tsxmap tsxTileSet

	err = xml.Unmarshal(byteValue, &tsxmap)
	if err != nil {
		panic(err)
	}

	return parseTileSet(&tsxmap)
}

// LoadTmx по файлу возвращаются
// layers - карта width x height
// tsxFileName - имя файла описания
func LoadTmx(filename string) (mymap *defines.Map, tsetsPtr *[]TileSetInfo) {
	fmt.Println("Loading map...", filename)

	xmlFile, err := resource.OpenFile(filename)
	if err != nil {
		panic(err)
	}
	defer resource.CloseFile(xmlFile)

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var tmxmap tmxMap

	err = xml.Unmarshal(byteValue, &tmxmap)
	if err != nil {
		panic(err)
	}

	lenLayers := len(tmxmap.Layers)
	layers := defines.Layers(make([]defines.Layer, lenLayers))

	for i, layer := range tmxmap.Layers {
		curlayer := parseLayer(layer)
		layers[i] = *curlayer
	}

	tilesets := make([]TileSetInfo, len(tmxmap.TileSets))

	for i, tileset := range tmxmap.TileSets {
		tilesets[i] = parseTileSet(tileset)
		fmt.Println("tileset name:", tileset.Name)
	}

	mymap = &defines.Map{
		Layers: layers,
		W:      layers[0].W, // ориентируемся по первому слою ( земля )
		H:      layers[0].H,
	}
	tsetsPtr = &tilesets

	return
}
