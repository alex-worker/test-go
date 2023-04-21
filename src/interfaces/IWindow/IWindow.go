package IWindow

import (
	. "test-go/src/defines"
)

type IWindow interface {
	GetInput() GameEvent
	DrawStart()
	DrawEnd()
}
