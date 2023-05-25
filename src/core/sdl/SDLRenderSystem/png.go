package SDLRenderSystem

import (
	"bytes"
	"github.com/veandco/go-sdl2/sdl"
	"image/png"
	"test-go/src/defines"
	. "test-go/src/math"
	"unsafe"
)

func PngBufToTexture(renderer *sdl.Renderer, buf *[]byte) (*sdl.Texture, Size2D, error) {
	bufReader := bytes.NewReader(*buf)

	myImage, err := png.Decode(bufReader)
	if err != nil {
		panic(err)
	}

	w := myImage.Bounds().Max.X
	h := myImage.Bounds().Max.Y

	pixels := make([]byte, w*h*4)
	bIndex := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, a := myImage.At(x, y).RGBA()
			pixels[bIndex] = byte(r / 256)
			bIndex++
			pixels[bIndex] = byte(g / 256)
			bIndex++
			pixels[bIndex] = byte(b / 256)
			bIndex++
			pixels[bIndex] = byte(a / 256)
			bIndex++
		}
	}

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, int32(w), int32(h))
	if err != nil {
		panic(err)
	}

	err = texture.Update(nil, unsafe.Pointer(&pixels[0]), w*4)
	if err != nil {
		panic(err)
	}

	//err = texture.SetBlendMode(sdl.BLENDMODE_BLEND)
	//if err != nil {
	//	panic(err)
	//}

	return texture, Size2D{Width: defines.Dimension(w), Height: defines.Dimension(h)}, nil
}
