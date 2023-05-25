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

type SDLViewMap2D struct {
	layer   []Cell
	size    Size2D
	tsx     *TileSet
	texture *SDLTexture
}

func New(size Size2D, tsx *TileSet, texture *SDLTexture) (*SDLViewMap2D, error) {
	if texture == nil {
		return nil, errors.New("invalid pointer")
	}

	layer := make([]Cell, size.Width*size.Height)

	return &SDLViewMap2D{
		layer:   layer,
		texture: texture,
		size:    size,
		tsx:     tsx,
	}, nil
}

func (m *SDLViewMap2D) Draw(r *SDLRenderSystem) error {

	var srcRect = sdl.Rect{
		X: 0,
		Y: 0,
		W: 300,
		H: 300,
	}

	var dstRect = sdl.Rect{
		X: 0,
		Y: 0,
		W: 300,
		H: 300,
	}

	var posX Dimension = 0
	var posY Dimension = 0

	for _ = range m.layer {

		err := r.GetRenderer().Copy(m.texture.Texture, &srcRect, &dstRect)
		if err != nil {
			return err
		}

		posX++
		if posX == m.size.Width {
			posX = 0
			posY++
		}
	}

	return nil
}
