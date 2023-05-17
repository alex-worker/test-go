package drawables

import (
	. "test-go/src/math"
)

type SDLViewMap2D struct {
	layer *Array2D
}

func (s *SDLViewMap2D) SetLayer(l *Array2D) {
	s.layer = l
}
