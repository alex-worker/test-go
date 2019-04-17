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
	263,263,263,263,263,263,263,263,263,263,
	263,263,263,263,263,263,263,263,263,263,
	263,263,263,263,263,263,263,263,263,263,
	263,263,263,263,263,263,263,263,263,263,
	263,263,263,263,263,263,263,263,263,263,
	263,263,263,263,263,263,263,263,263,263,
	263,263,263,263,263,263,263,263,263,263,
	263,263,263,263,263,263,263,263,263,263,
	263,263,263,263,263,263,263,263,263,263,
	263,263,263,263,263,263,263,263,263,263,
}

var mockCells2 = Layer{
	815,815,815,815,815,815,815,815,815,815,
	815,815,815,815,815,815,815,815,815,815,
	815,815,815,815,815,815,815,815,815,815,
	815,815,815,815,815,815,815,815,815,815,
	815,815,815,815,815,815,815,815,815,815,
	815,815,815,815,815,815,815,815,815,815,
	815,815,815,815,815,815,815,815,815,815,
	815,815,815,815,815,815,815,815,815,815,
	815,815,815,815,815,815,815,815,815,815,
	815,815,815,815,815,815,815,815,815,815,
}

var mockLayers = Layers{
	&mockCells,
	&mockCells2,
}

// MakeView делаем окно по карте
func (v *View) MakeView(m *def.Map, pos def.Pos, sz def.Size) {
	v.pos = pos
	v.Size = sz
	v.Layers = make(Layers, len(m.Layers) )
	for i:=0;i<len(m.Layers);i++{
		v.Layers[i] = v.importLayer( &(m.Layers[i]), pos )
	}
	
}

// импортнуть один уровень с карты
func (v *View) importLayer( src *def.Layer, pos def.Pos) *Layer {

	var x,y uint32
	index := 0

	layer := make(Layer, v.Size.Width*v.Size.Height )

	for y=pos.Y; y<pos.Y+v.Size.Height; y++ {
		for x=pos.X; x<pos.X+v.Size.Width; x++ {
			cell := (*src.Data)[y][x]
			layer[index] = cell
			index++
		}
	}

	return &layer

}

// GetSize размер окошка
func (v* View) GetSize() def.Size {
	return v.Size
}

// GetView получаем массив просматриваемых тайлов
func (v* View) GetView() *Layers {
	return &v.Layers
}
