package Tiles

import (
	. "test-go/src/interfaces/IWindow"
)

type DrawableTileMap struct {
}

func (d *DrawableTileMap) Init(window *IWindow) {
	//TODO implement me
	panic("implement me")
}

func (d *DrawableTileMap) Draw(window *IWindow) {
	//TODO implement me
	//panic("implement me")
}

func GetDrawableTileMap() IDrawable {
	item := DrawableTileMap{}
	return &item
}
