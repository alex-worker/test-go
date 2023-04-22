package TileMap

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	. "test-go/src/core/TileMap/parser"
)

func strToUint(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}

func convertLayer(layer *TmxLayer) (*Layer, error) {
	fmt.Printf("layer data: %#v %#v\n", layer.Width, layer.Height)

	re := regexp.MustCompile(`\r?\n`)
	normalizedMap := re.ReplaceAllString(layer.Data, "")
	myMapStr := strings.Split(normalizedMap, ",")

	w, err := strToUint(layer.Width)
	if err != nil {
		return nil, err
	}

	h, err := strToUint(layer.Height)
	if err != nil {
		return nil, err
	}

	cells := make([]Cell, w*h)

	var index uint64
	for _, c := range myMapStr {
		cell, err := strToUint(c)
		if err != nil {
			panic(err)
		}
		cells[index] = Cell(cell)
		index++
	}

	return &Layer{
		Data: cells,
		W:    w,
		H:    h,
		Name: layer.Name,
	}, nil
}

func convertTileSet(set *TsxTileSet) *TileSetInfo {
	var tsxFileName = set.Source
	if tsxFileName != "" {
		panic("can't implement file download")
	}

	w, err := strToUint(set.TileWidth)
	if err != nil {
		panic(err)
	}

	h, err := strToUint(set.TileHeight)
	if err != nil {
		panic(err)
	}

	fileName := set.Image.Source

	lenTiles := len(set.Tiles)
	fmt.Printf("Tiles: %#v\n", lenTiles)
	tiles := make([]AnimateTile, lenTiles)

	for _, tile := range set.Tiles {
		fmt.Printf("tile: #%v\n", tile)
		fmt.Println("name:", tile.XMLName)
		fmt.Println("ID:", tile.ID)
	}

	return &TileSetInfo{
		ImageFileName: fileName,
		Tiles:         tiles,
		TileW:         w,
		TileH:         h,
	}
}

func convertAnimateTiles(set *TsxTileSet) *[]AnimateTile {

	return nil
}
