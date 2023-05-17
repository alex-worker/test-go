package SDLRenderSystem

import (
	. "test-go/src/core/sdl"
	. "test-go/src/interfaces/IRenderSystem"
	. "test-go/src/math"
)

func GetRenderSystem(windowSize Size2D) (IRenderSystem, error) {
	renderer := InitSDL(windowSize)
	renderSystem := SDLRenderSystem{
		renderer: renderer,
	}
	return &renderSystem, nil
}
