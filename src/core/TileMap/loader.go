package TileMap

import (
	"fmt"
	"strconv"
	. "test-go/src/core/TileMap/parser"
)

func Load(m *TmxMap) (*TileMap, error) {
	lenLayers := len(m.Layers)
	fmt.Printf("layers: %#v\n", lenLayers)

	layers := make([]Layer, lenLayers)
	for i, layer := range m.Layers {
		curlayer, err := parseLayer(layer)
		if err != nil {
			return nil, err
		}
		layers[i] = *curlayer
	}

	fmt.Println(layers)
	return nil, nil
}

func parseLayer(layer *TmxLayer) (*Layer, error) {
	fmt.Printf("layer data: %#v %#v\n", layer.Width, layer.Height)
	w, err := readUint(layer.Width)
	if err != nil {
		return nil, err
	}
	h, err := readUint(layer.Height)
	if err != nil {
		return nil, err
	}

	return &Layer{
		Data: nil,
		W:    w,
		H:    h,
		Name: layer.Name,
	}, nil
}

func readUint(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}
