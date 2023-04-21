package TileMap

import (
	"fmt"
	. "test-go/src/core/TileMap/parser"
)

func Load(m *TmxMap) (*TileMap, error) {
	lenLayers := len(m.Layers)
	fmt.Printf("layers: %#v\n", lenLayers)

	layers := make([]Layer, lenLayers)
	for i, tmxLayer := range m.Layers {
		curLayer, err := convertLayer(tmxLayer)
		if err != nil {
			return nil, err
		}
		layers[i] = *curLayer
	}

	fmt.Printf("tilesets: %#v\n", m.TileSets[0])
	return &TileMap{
		Layers: &layers,
		W:      0,
		H:      0,
	}, nil
}
