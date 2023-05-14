package IRenderSystem

import (
	. "test-go/src/defines"
)

type IRenderSystem interface {
	GetInput() GameEvent
	Draw()
}
