package SDLViewMap2D

import (
	"errors"
	"github.com/veandco/go-sdl2/sdl"
	. "test-go/src/core/TileMap/TileAnimations"
	. "test-go/src/core/sdl/SDLRenderSystem"
	. "test-go/src/core/sdl/SDLTexture"
	. "test-go/src/defines"
	. "test-go/src/math"
	Array2D "test-go/src/math/Array2D"
)

type SizeInt32 struct {
	Width  int32
	Height int32
}

type PosInt32 struct {
	X int32
	Y int32
}

type SDLViewMap2D struct {
	a         Array2D.Array2D
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

	arr := Array2D.New(size)

	tileShift := calcTileShift(tsx)

	return &SDLViewMap2D{
		a: arr,
		//layer:     layer,
		texture: texture,
		//size:      size,
		tsx:       tsx,
		tileShift: tileShift,
	}, nil
}

func (m *SDLViewMap2D) Draw(r *SDLRenderSystem) error {

	screenSize := r.GetScreenSize()

	scrDeltaX := int32(screenSize.Width / m.a.Size.Width)
	scrDeltaY := int32(screenSize.Height / m.a.Size.Height)
	if scrDeltaX > scrDeltaY {
		scrDeltaX = scrDeltaY
	} else {
		scrDeltaY = scrDeltaX
	}

	//fmt.Println("scrDeltaX: ", scrDeltaX, "scrDeltaY: ", scrDeltaY)

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

	for range m.a.Data {

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

		if posX == m.a.Size.Width {
			posX = 0
			posY++
			scrPosX = 0
			scrPosY += scrDeltaY
		}
	}

	return nil
}

func (m *SDLViewMap2D) GetFromMap(pos Point2D, src *Array2D.Array2D) error {
	return nil
}

func (m *SDLViewMap2D) Update(delta uint64) error {
	return nil
}

func (m *SDLViewMap2D) tileToRect(tile Cell, pos *sdl.Rect) {
	cY := int32(tile) >> m.tileShift
	cX := int32(tile) - cY<<m.tileShift

	pos.X = cY * int32(m.tsx.TileW)
	pos.Y = cX * int32(m.tsx.TileH)
}

func calcTileShift(t *TileSet) uint32 {
	var tileShift uint32 = 1

	tilesCnt := t.Columns

	for tilesCnt > 2 { // ручной логарифм по основанию 2 !
		tilesCnt = tilesCnt / 2
		tileShift++
	}
	return tileShift
}
