package IWindow

import (
	. "test-go/src/defines"
)

type ITexture interface{}

type IWindow interface {
	GetInput() GameEvent
	DrawStart()
	DrawEnd()
	DecodeTexture(buf *[]byte) (ITexture, error)
}

type IDrawable interface {
	Init(window *IWindow)
	Draw(window *IWindow)
}
