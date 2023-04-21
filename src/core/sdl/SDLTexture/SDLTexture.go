package SDLTexture

import (
	"github.com/veandco/go-sdl2/sdl"
	"test-go/src/defines"
)

type SDLTexture struct {
	Texture *sdl.Texture
	Size    defines.Size
}
