package SDLRenderSystem

import (
	"github.com/veandco/go-sdl2/sdl"
	. "test-go/src/core/sdl"
	. "test-go/src/core/sdl/SDLTexture"
	. "test-go/src/math"
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
	startTicks := sdl.GetTicks64()
	s.drawStart()
	s.drawEnd()
	endTicks := sdl.GetTicks64()
	fps := CalcFPS(startTicks, endTicks)
	println(fps)
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
