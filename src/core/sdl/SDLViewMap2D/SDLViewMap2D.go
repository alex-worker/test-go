package SDLViewMap2D

import (
	"errors"
	"github.com/veandco/go-sdl2/sdl"
	. "test-go/src/core/TileMap/TileAnimations"
	. "test-go/src/core/sdl/SDLRenderSystem"
	. "test-go/src/core/sdl/SDLTexture"
	. "test-go/src/defines"
	. "test-go/src/math"
)

type PosInt32 struct {
	X int32
	Y int32
}

type SDLViewMap2D struct {
	layer     []Cell
	size      Size2D
	tsx       *TileSet
	tileShift uint32
	texture   *SDLTexture
}

func New(size Size2D, tsx *TileSet, texture *SDLTexture) (*SDLViewMap2D, error) {
	if texture == nil {
		return nil, errors.New("invalid pointer")
	}
	if tsx == nil {
		return nil, errors.New("invalid pointer")
	}

	layer := make([]Cell, size.Width*size.Height)

	tileShift := CalcTileShift(tsx)

	return &SDLViewMap2D{
		layer:     layer,
		texture:   texture,
		size:      size,
		tsx:       tsx,
		tileShift: tileShift,
	}, nil
}

func (m *SDLViewMap2D) Draw(r *SDLRenderSystem) error {

	// условно считаем что тайлы квадратные и у массива тайлов тоже длина и высота одна и та же
	screenSize := r.GetScreenSize()
	minScreenSize := Min(screenSize.Width, screenSize.Height)

	//minViewSize := Min(m.size.Width, m.size.Height)

	var scrDeltaX = int32(minScreenSize / m.size.Width)
	var scrDeltaY = int32(minScreenSize / m.size.Height)

	// дальше пропорция:
	// minScreenSize -> m.size.Width

	//var scrDeltaX = int32(m.tsx.TileW)
	//var scrDeltaY = int32(m.tsx.TileH)

	srcRect := sdl.Rect{
		X: 0,
		Y: 0,
		W: int32(m.tsx.TileW),
		H: int32(m.tsx.TileH),
	}

	dstRect := sdl.Rect{
		X: 0,
		Y: 0,
		W: scrDeltaX,
		H: scrDeltaY,
	}

	var scrPosX int32 = 0
	var scrPosY int32 = 0

	var posX Dimension = 0
	var posY Dimension = 0

	for _ = range m.layer {

		//m.tileToRect(Cell(index), &srcRect)
		srcRect.X = int32(posX) * int32(m.tsx.TileW)
		srcRect.Y = int32(posY) * int32(m.tsx.TileH)

		dstRect.X = scrPosX
		dstRect.Y = scrPosY

		err := r.GetRenderer().Copy(m.texture.Texture, &srcRect, &dstRect)
		if err != nil {
			return err
		}

		posX++
		scrPosX += scrDeltaX

		if posX == m.size.Width {
			posX = 0
			posY++
			scrPosX = 0
			scrPosY += scrDeltaY
		}
	}

	return nil
}

func (m *SDLViewMap2D) tileToRect(tile Cell, pos *sdl.Rect) {
	//if tile == 0 {
	//	tile = 12
	//}
	cY := int32(tile) >> m.tileShift
	cX := int32(tile) - cY<<m.tileShift

	pos.X = cY * int32(m.tsx.TileW)
	pos.Y = cX * int32(m.tsx.TileH)
}

func CalcTileShift(t *TileSet) uint32 {
	var tileShift uint32 = 1

	tilesCnt := t.Columns

	for tilesCnt > 2 { // ручной логарифм по основанию 2 !
		tilesCnt = tilesCnt / 2
		tileShift++
	}
	return tileShift
}
