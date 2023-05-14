package IInputSystem

import (
	. "test-go/src/defines"
)

type IInputSystem interface {
	GetInput() GameEvent
}
