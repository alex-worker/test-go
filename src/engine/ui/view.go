package ui

import (
	"../def"
)

// Layer Один слой
type Layer []def.Cell // одномерный массив тайлов 

// Layers Несколько слоев
type Layers []*Layer

// View Окно обзора
type View struct {
	Layers Layers
	pos def.Pos
	Size def.Size
}

var mockCells = Layer{
	1,1,1,1,1,1,1,1,1,1,
	1,0,0,0,0,0,0,0,0,1,
	1,0,0,0,0,0,0,0,0,1,
	1,0,0,0,0,0,0,0,0,1,
	1,0,0,0,0,0,0,0,0,1,
	1,0,0,0,0,0,0,0,0,1,
	1,0,0,0,0,0,0,0,0,1,
	1,0,0,0,0,0,0,0,0,1,
	1,0,0,0,0,0,0,0,0,1,
	1,1,1,1,1,1,1,1,1,1,
}

var mockCells2 = Layer{
	5,5,5,5,5,5,5,5,5,5,
	5,5,5,5,5,5,5,5,5,5,
	5,5,5,5,5,5,5,5,5,5,
	5,5,5,5,5,5,5,5,5,5,
	5,5,5,5,5,5,5,5,5,5,
	5,5,5,5,5,5,5,5,5,5,
	5,5,5,5,5,5,5,5,5,5,
	5,5,5,5,5,5,5,5,5,5,
	5,5,5,5,5,5,5,5,5,5,
	5,5,5,5,5,5,5,5,5,5,
}

// var mockLayer = Layer(mockCells)

var mockLayers = Layers{
	&mockCells,
	&mockCells2,
}

// MakeView делаем окно по карте
func (v *View) MakeView(Map *def.Map, pos def.Pos, sz def.Size) {
	v.pos = pos
	v.Size = sz
	v.Layers = mockLayers
}

// GetSize размер окошка
func (v* View) GetSize() def.Size {
	return v.Size
}

// GetView получаем массив просматриваемых тайлов
func (v* View) GetView() *Layers {
	return &v.Layers
}
