package SDLTexture

import (
	"github.com/veandco/go-sdl2/sdl"
	"test-go/src/math"
)

type SDLTexture struct {
	Size    math.Size2D
	Texture *sdl.Texture
}
