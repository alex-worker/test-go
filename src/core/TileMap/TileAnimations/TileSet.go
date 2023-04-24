package TileAnimations

import "test-go/src/core/TileMap"

// TileSet описание тайлсета для дальнейшей обработки
type TileSet struct {
	//	ImageFileName string
	Tiles map[TileMap.Cell]AnimateTile
	TileW uint64
	TileH uint64
}
