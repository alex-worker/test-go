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

	lenTileSets := len(m.TileSets)
	fmt.Printf("TileSets: %#v\n", lenTileSets)
	tileSets := make([]*TileSetInfo, lenTileSets)
	fmt.Printf("tilesets: #%v\n", tileSets)

	for i, tileset := range m.TileSets {
		tileSets[i] = convertTileSet(tileset)
		fmt.Println("tileset name:", tileset.Name)
	}

	return &TileMap{
		Layers: &layers,
		W:      layers[0].W, // ориентируемся по первому слою (земля)
		H:      layers[0].H,
	}, nil
}

func LoadTilesets(m *TmxMap) (*[]TileSetInfo, error) {
	return nil, nil
}
