package IRenderSystem

import (
	. "test-go/src/defines"
)

type IViewMap2D interface {
	GetCell(p Point2D) Cell
	SetCell(p Point2D, c Cell) error
}
