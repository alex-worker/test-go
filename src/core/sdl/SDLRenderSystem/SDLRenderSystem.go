package SDLRenderSystem

import (
	"github.com/veandco/go-sdl2/sdl"
	. "test-go/src/core/sdl"
	"test-go/src/defines"
	. "test-go/src/interfaces/IRenderSystem"
)

type SDLRenderSystem struct {
	renderer *sdl.Renderer
}

func (s *SDLRenderSystem) drawStart() {
	err := s.renderer.Clear()
	if err != nil {
		panic(err)
	}
}

func (s *SDLRenderSystem) drawEnd() {
	s.renderer.Present()
}

func (s *SDLRenderSystem) Draw() {
	s.drawStart()
	s.drawEnd()
}

func GetRenderSystem(windowSize defines.Size) (IRenderSystem, error) {
	renderer := InitSDL(windowSize)
	renderSystem := SDLRenderSystem{
		renderer: renderer,
	}
	return &renderSystem, nil
}
