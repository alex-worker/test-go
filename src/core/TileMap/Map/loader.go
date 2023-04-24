package Map

import (
	"fmt"
	"regexp"
	"strings"
	TileMap2 "test-go/src/core/TileMap"
	. "test-go/src/core/TileMap/parser"
)

func LoadMap(m *TmxMap) (*TileMap, error) {
	lenLayers := len(m.Layers)
	fmt.Printf("Layers: %#v\n", lenLayers)

	layers := make([]Layer, lenLayers)
	for i, tmxLayer := range m.Layers {
		curLayer, err := convertLayer(tmxLayer)
		if err != nil {
			return nil, err
		}
		layers[i] = *curLayer
	}

	return &TileMap{
		Layers: layers,
		W:      layers[0].W, // ориентируемся по первому слою (земля)
		H:      layers[0].H,
	}, nil
}

func convertLayer(layer *TmxLayer) (*Layer, error) {
	fmt.Printf("layer data: %#v %#v\n", layer.Width, layer.Height)

	re := regexp.MustCompile(`\r?\n`)
	normalizedMap := re.ReplaceAllString(layer.Data, "")
	myMapStr := strings.Split(normalizedMap, ",")

	w, err := TileMap2.StrToUint(layer.Width)
	if err != nil {
		return nil, err
	}

	h, err := TileMap2.StrToUint(layer.Height)
	if err != nil {
		return nil, err
	}

	cells := make([]TileMap2.Cell, w*h)

	var index uint64
	for _, c := range myMapStr {
		cell, err2 := TileMap2.StrToUint(c)
		if err2 != nil {
			panic(err2)
		}
		cells[index] = TileMap2.Cell(cell)
		index++
	}

	return &Layer{
		Data: cells,
		W:    w,
		H:    h,
		Name: layer.Name,
	}, nil
}
