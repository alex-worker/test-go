package SDLRenderSystem

import (
	a "test-go/src/math/Array2D"
)

type SDLViewMap2D struct {
	layer *a.Array2D
}

func New(l *a.Array2D) SDLViewMap2D {
	return SDLViewMap2D{
		layer: l,
	}
}

func (s *SDLViewMap2D) SetLayer(l *a.Array2D) {
	s.layer = l
}
