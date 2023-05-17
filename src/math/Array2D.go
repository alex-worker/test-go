package math

import (
	. "test-go/src/defines"
)

type Array2D struct {
	data []Cell
	size Size2D
}

func New(s Size2D) Array2D {
	data := make([]Cell, s.Width*s.Height)
	return Array2D{
		data: data,
		size: s,
	}
}

func (a *Array2D) GetCell(p Point2D) (Cell, error) {
	index, err := a.size.GetIndex(p)
	if err != nil {
		return 0, err
	}
	return a.data[index], nil
}
