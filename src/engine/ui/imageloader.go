package ui

import (
	"os"
	"image/png"
	"github.com/veandco/go-sdl2/sdl"
)

func imgFileToTexture(filename string) *sdl.Texture {
	infile, err := os.Open( filename )
	if err != nil {
		panic(err)
	}

	img, err := png.Decode(infile)
	if err != nil {
		panic(err)
	}

	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y

	pixels := make([]byte,w*h*4)
	bIndex := 0
	for y:=0; y < h; y++ {
		for x:=0; x < w; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			pixels[bIndex] = byte(r / 256 )
			bIndex++
			pixels[bIndex] = byte(g / 256 )
			bIndex++
			pixels[bIndex] = byte(b / 256 )
			bIndex++
			pixels[bIndex] = byte(a / 256 )
			bIndex++
		}
	}

	tex,err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, int32(w), int32(h) )
	if err != nil {
		panic(err)
	}
	tex.Update(nil, pixels, w*4)

	err = tex.SetBlendMode(sdl.BLENDMODE_BLEND)
	if err != nil {
		panic(err)
	}

	return tex
}