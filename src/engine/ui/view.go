package ui

import (
	"../def"
)

// Layer Один слой
type Layer []def.Cell // одномерный массив тайлов 

// Layers Несколько слоев
type Layers []*Layers

// View Окно обзора
type View struct {
	Map *def.Map
	Layers Layers
	pos def.Pos
}

// MakeView делаем окно по карте
func (v *View) MakeView(pos def.Pos, sz def.Size) {
	v.pos = pos
}

// SetMap ставим карту
func (v *View) SetMap( mymap *def.Map) {
	v.Map = mymap
}

// GetView получаем массив просматриваемых тайлов
func (v* View) GetView() *Layers {
	return &v.Layers
}
