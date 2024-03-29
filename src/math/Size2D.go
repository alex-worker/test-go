package math

import (
	"errors"
	. "test-go/src/defines"
)

// Size2D длина и ширина
type Size2D struct {
	Width  Dimension
	Height Dimension
}

func (s Size2D) IsPointInto(p Point2D) bool {
	return s.Width > p.X && s.Height > p.Y // (!) without >=
}

func (s Size2D) GetIndex(p Point2D) (Dimension, error) {
	if s.IsPointInto(p) {
		return s.Width*p.Y + p.X, nil
	}
	return 0, errors.New("out of size")
}
