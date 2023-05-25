package SDLRenderSystem

import (
	"github.com/veandco/go-sdl2/sdl"
	. "test-go/src/core/sdl/SDLTexture"
	. "test-go/src/math"
)

type SDLRenderSystem struct {
	renderer   *sdl.Renderer
	windowSize Size2D
	ticks      uint64
	fps        uint64
}

func (s *SDLRenderSystem) GetScreenSize() Size2D {
	return s.windowSize
}

func (s *SDLRenderSystem) GetFPS() uint64 {
	return s.fps
}

func (s *SDLRenderSystem) GetRenderer() *sdl.Renderer {
	return s.renderer
}

func (s *SDLRenderSystem) DrawStart() error {
	s.ticks = sdl.GetTicks64()
	return s.renderer.Clear()
}

func (s *SDLRenderSystem) DrawEnd() {
	s.renderer.Present()
	endTicks := sdl.GetTicks64()
	s.fps = CalcFPS(s.ticks, endTicks)
}

func (s *SDLRenderSystem) GetTexture(buf *[]byte) (*SDLTexture, error) {
	texture, size, err := PngBufToTexture(s.renderer, buf)
	if err != nil {
		return nil, err
	}
	return &SDLTexture{
		Size:    size,
		Texture: texture,
	}, nil
}
