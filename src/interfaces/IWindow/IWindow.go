package IWindow

import (
	. "test-go/src/defines"
)

type ITexture interface{}

type IWindow interface {
	GetInput() GameEvent
	DrawStart()
	DrawEnd()
}

type IDrawable interface {
	Init(window *IWindow)
	Draw(window *IWindow)
}
