package Tiles

import (
	. "test-go/src/core/sdl/SDLTexture"
	. "test-go/src/interfaces/IWindow"
)

type DrawableTileMap struct {
	texture *SDLTexture
}

func (d *DrawableTileMap) Init(window *IWindow) {
	//TODO implement me
	panic("implement me")
}

func (d *DrawableTileMap) Draw(window *IWindow) {
	//TODO implement me
	//panic("implement me")
}

//func GetDrawableTileMap(tex *ITexture) IDrawable {
//	texture := SDLTexture(*tex)
//
//	item := DrawableTileMap{
//		texture: texture,
//	}
//	return &item
//}
