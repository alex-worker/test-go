package SDLTexture

import (
	"github.com/veandco/go-sdl2/sdl"
	"test-go/src/math"
)

func GetSDLTexture(size math.Size2D, texture *sdl.Texture) (*SDLTexture, error) {
	return &SDLTexture{
		Size:    size,
		Texture: texture,
	}, nil
}
