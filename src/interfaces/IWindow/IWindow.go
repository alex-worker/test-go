package IWindow

import (
	. "test-go/src/defines"
)

type IWindow interface {
	GetInput() GameEvent
	DrawStart()
	Draw(item *IDrawable)
	DrawEnd()
}

type IDrawable interface {
	Init(window *IWindow)
	Draw(window *IWindow)
}
