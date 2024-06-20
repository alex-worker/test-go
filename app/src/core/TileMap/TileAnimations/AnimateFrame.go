package TileAnimations

import (
	. "test-go/src/defines"
)

// AnimateFrame анимационный фрейм тайла
type AnimateFrame struct {
	Cell     Cell   // номер тайла
	Duration uint32 // задержка таймера
}
