package SDLRenderSystem

import (
	. "test-go/src/core/sdl"
	. "test-go/src/defines"
	. "test-go/src/interfaces/IRenderSystem"
)

func GetRenderSystem(windowSize Size2D) (IRenderSystem, error) {
	renderer := InitSDL(windowSize)
	renderSystem := SDLRenderSystem{
		renderer: renderer,
	}
	return &renderSystem, nil
}
