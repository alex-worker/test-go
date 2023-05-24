package SDLViewMap2D

import (
	"errors"
	"github.com/veandco/go-sdl2/sdl"
	. "test-go/src/core/sdl/SDLRenderSystem"
	. "test-go/src/math"
	"test-go/src/math/Array2D"
)

type SDLViewMap2D struct {
	layer   Array2D.Array2D
	texture *sdl.Texture
}

func New(size Size2D, texture *sdl.Texture) (*SDLViewMap2D, error) {
	layer := Array2D.New(size)
	if texture == nil {
		return nil, errors.New("invalid pointer")
	}
	return &SDLViewMap2D{
		layer:   layer,
		texture: texture,
	}, nil
}

func (m *SDLViewMap2D) Draw(r *SDLRenderSystem) error {
	srcRect := sdl.Rect{
		X: 0,
		Y: 0,
		W: 300,
		H: 300,
	}

	dstRect := sdl.Rect{
		X: 0,
		Y: 0,
		W: 300,
		H: 300,
	}

	return r.GetRenderer().Copy(m.texture, &srcRect, &dstRect)

}
