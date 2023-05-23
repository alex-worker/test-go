package SDLViewMap2D

import (
	"errors"
	a "test-go/src/math/Array2D"
)

type SDLViewMap2D struct {
	layer *a.Array2D
}

func New(l *a.Array2D) (*SDLViewMap2D, error) {
	if l == nil {
		return nil, errors.New("invalid pointer")
	}
	return &SDLViewMap2D{
		layer: l,
	}, nil
}
