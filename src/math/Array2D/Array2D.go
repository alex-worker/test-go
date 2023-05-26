package Array2D

import (
	. "test-go/src/defines"
	. "test-go/src/math"
)

type Array2D struct {
	Data []Cell
	size Size2D
}

func New(s Size2D) Array2D {
	data := make([]Cell, s.Width*s.Height)
	return Array2D{
		Data: data,
		size: s,
	}
}

func (a *Array2D) GetCell(p Point2D) (Cell, error) {
	index, err := a.size.GetIndex(p)
	if err != nil {
		return 0, err
	}
	return a.Data[index], nil
}
