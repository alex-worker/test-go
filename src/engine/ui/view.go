package ui

import (
	"test-go/src/engine/def"
	. "test-go/src/math"
)

// Layer Один слой
type Layer []def.Cell // одномерный массив тайлов

// Layers Несколько слоев
type Layers []*Layer

// View Окно обзора
type View struct {
	Layers Layers
	pos    Pos2D
	Size   Size2D
}

// MakeView делаем окно по карте
func (v *View) MakeView(m *def.Map, pos Pos2D, sz Size2D) {
	v.pos = pos
	v.Size = sz

	maxPosX := m.W - sz.Width - 1
	maxPosY := m.H - sz.Height - 1

	if v.pos.X > maxPosX {
		v.pos.X = maxPosX
	}

	if v.pos.Y > maxPosY {
		v.pos.Y = maxPosY
	}

	v.Layers = make(Layers, len(m.Layers))
	for i := 0; i < len(m.Layers); i++ {
		v.Layers[i] = v.importLayer(&(m.Layers[i]), v.pos)
	}

}

// импортнуть один уровень с карты
func (v *View) importLayer(src *def.Layer, pos Pos2D) *Layer {

	var x, y uint32
	index := 0

	layer := make(Layer, v.Size.Width*v.Size.Height)

	for y = pos.Y; y < pos.Y+v.Size.Height; y++ {
		for x = pos.X; x < pos.X+v.Size.Width; x++ {
			cell := (*src.Data)[y][x]
			layer[index] = cell
			index++
		}
	}

	return &layer
}

// GetSize размер окошка
func (v *View) GetSize() Size2D {
	return v.Size
}

// GetView получаем массив просматриваемых тайлов
func (v *View) GetView() *Layers {
	return &v.Layers
}
