package SDLRenderSystem

import (
	. "test-go/src/math"
)

func GetRenderSystem(windowSize Size2D) (*SDLRenderSystem, error) {
	renderer := InitSDL(windowSize)
	renderSystem := SDLRenderSystem{
		renderer: renderer,
	}
	return &renderSystem, nil
}
