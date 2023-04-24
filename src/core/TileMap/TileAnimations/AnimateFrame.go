package TileAnimations

import "test-go/src/core/TileMap"

// AnimateFrame анимационный фрейм тайла
type AnimateFrame struct {
	Cell     TileMap.Cell // номер тайла
	Duration uint32       // задержка таймера
}
