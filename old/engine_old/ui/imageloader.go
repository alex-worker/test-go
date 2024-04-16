package ui

import (
	"bytes"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"image/png"
	"io"
	"test-go/src/engine/resource"
	"unsafe"
)

func imgFileToTexture(filename string) (texture *sdl.Texture, w int, h int) {
	fmt.Println("Load texture...", filename)
	infile, err := resource.OpenFile(filename)
	if err != nil {
		panic(err)
	}
	buf, err := io.ReadAll(infile)
	if err != nil {
		panic(err)
	}

	bufReader := bytes.NewReader(buf)

	img, err := png.Decode(bufReader)
	if err != nil {
		panic(err)
	}

	w = img.Bounds().Max.X
	h = img.Bounds().Max.Y

	pixels := make([]byte, w*h*4)
	bIndex := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, a := img.At(x, y).RGBA()
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

	texture, err = renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, int32(w), int32(h))
	if err != nil {
		panic(err)
	}

	err = texture.Update(nil, unsafe.Pointer(&pixels[0]), w*4)
	if err != nil {
		panic(err)
	}

	err = texture.SetBlendMode(sdl.BLENDMODE_BLEND)
	if err != nil {
		panic(err)
	}
	return
}