package TileMap

import (
	"fmt"
	"strconv"
	. "test-go/src/core/TileMap/parser"
)

func convertLayer(layer *TmxLayer) (*Layer, error) {
	fmt.Printf("layer data: %#v %#v\n", layer.Width, layer.Height)

	w, err := strToUint(layer.Width)
	if err != nil {
		return nil, err
	}

	h, err := strToUint(layer.Height)
	if err != nil {
		return nil, err
	}

	cells := make([]Cell, w*h)

	return &Layer{
		Data: &cells,
		W:    w,
		H:    h,
		Name: layer.Name,
	}, nil
}

func strToUint(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}
