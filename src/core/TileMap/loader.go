package TileMap

import (
	"fmt"
	. "test-go/src/core/TileMap/parser"
)

func Load(m *TmxMap) (*TileMap, error) {
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

	lenTileSets := len(m.TileSets)
	fmt.Printf("TileSets: %#v\n", lenTileSets)
	//tilesets := make([]TileSetInfo, len(tmxmap.TileSets))

	return &TileMap{
		Layers: &layers,
		W:      0,
		H:      0,
	}, nil
}
