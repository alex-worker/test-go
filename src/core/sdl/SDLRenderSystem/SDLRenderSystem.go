package SDLRenderSystem

import (
	"github.com/veandco/go-sdl2/sdl"
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
