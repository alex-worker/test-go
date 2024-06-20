package TileAnimations

import (
	. "test-go/src/defines"
)

// TileSet описание тайлсета для дальнейшей обработки
type TileSet struct {
	//	ImageFileName string
	Tiles     map[Cell]AnimateTile
	TileW     uint64
	TileH     uint64
	Columns   uint64
	TileCount uint64
}
