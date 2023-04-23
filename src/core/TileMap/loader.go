package TileMap

import (
	"fmt"
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

func LoadTileSets(m *TmxMap) ([]TileSet, error) {
	lenTileSets := len(m.TileSets)
	tileSets := make([]TileSet, lenTileSets)

	for i, tsxTileSet := range m.TileSets {
		curTileSet := convertTileSet(tsxTileSet)
		tileSets[i] = *curTileSet
	}
	return tileSets, nil
}
