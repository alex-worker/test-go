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
		cell, err2 := strToUint(c)
		if err2 != nil {
			panic(err2)
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

func convertFrames(anims *TsxAnimation) []AnimateFrame {
	lenFrames := len(anims.Frames)

	frames := make([]AnimateFrame, lenFrames)

	for i, f := range anims.Frames {
		frames[i] = AnimateFrame{
			Cell:     Cell(f.Tileid),
			Duration: f.Duration,
		}
	}
	return frames
}

func convertTileSet(set *TsxTileSet) *TileSet {
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

	tiles := make(map[Cell]AnimateTile)

	for _, tile := range set.Tiles {
		idCell := Cell(tile.ID)
		//fmt.Println("ID:", idCell)

		frames := convertFrames(&tile.Animations)

		tiles[idCell] = AnimateTile{
			Tick:       0,
			Index:      0,
			NeedUpdate: false,
			Frames:     frames,
		}
	}

	return &TileSet{
		ImageFileName: fileName,
		Tiles:         tiles,
		TileW:         w,
		TileH:         h,
	}
}
