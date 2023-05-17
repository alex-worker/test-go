package math

import (
	. "test-go/src/defines"
)

type Array2D struct {
	data []Cell
	size Size2D
}

func (a *Array2D) GetCell(p Point2D) (Cell, error) {
	index, err := a.size.GetIndex(p)
	if err != nil {
		return 0, err
	}
	return a.data[index], nil
}
