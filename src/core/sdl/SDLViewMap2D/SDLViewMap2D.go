package SDLViewMap2D

import (
	"errors"
	"github.com/veandco/go-sdl2/sdl"
	a "test-go/src/math/Array2D"
)

type SDLViewMap2D struct {
	layer   *a.Array2D
	texture *sdl.Texture
}

func New(l *a.Array2D, t *sdl.Texture) (*SDLViewMap2D, error) {
	if l == nil {
		return nil, errors.New("invalid pointer")
	}
	if t == nil {
		return nil, errors.New("invalid pointer")
	}
	return &SDLViewMap2D{
		layer:   l,
		texture: t,
	}, nil
}
