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

func (t *TileSet) CalcTileShift() uint32 {
	var tileShift uint32 = 1

	tilesCnt := t.Columns

	for tilesCnt > 2 { // ручной логарифм по основанию 2 !
		tilesCnt = tilesCnt / 2
		tileShift++
	}
	return tileShift
}
