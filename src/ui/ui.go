package ui

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"image/png"
	"../world"
)

var window *sdl.Window = nil
var renderer *sdl.Renderer = nil
var textureAtlas *sdl.Texture = nil

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

// return false when window is closed
func Update() bool {

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		default:
			// fmt.Println( event )
			break
		case *sdl.QuitEvent:
			println("Quit")
			return false
		}
	}

	return true
}

func Destroy(){
	renderer.Destroy()
	window.Destroy()
	sdl.Quit()
	fmt.Println("Ui offline...")
}

func Init(){

	// sdl.LogSetAllPriority(sdl.LOG_PRIORITY_VERBOSE)
	err := sdl.Init(sdl.INIT_EVERYTHING); 
	if err != nil {
		panic(err)
	}

	window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
    800, 600, sdl.WINDOW_SHOWN )
    if err != nil {
        panic(err)
    }
	
	// SDL_RENDERER_ACCELERATED для хардварной поддержки
	renderer, err = sdl.CreateRenderer( window, -1, sdl.RENDERER_SOFTWARE)
    if err != nil {
        panic(err)
	}

	// surface, err := window.GetSurface()
    // if err != nil {
        // panic(err)
    // }
	// surface.FillRect(nil, 0)
	// window.UpdateSurface()

	// sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	textureAtlas = imgFileToTexture("ui/assets/tiles.png")
	
	renderer.Copy(textureAtlas, nil, nil)
	renderer.Present()
	// surface.FillRect(nil, 0)
    // // rect := sdl.Rect{0, 0, 200, 200}
    // surface.FillRect(&rect, 0xffff0000)

}

type Ui struct {
}

func Draw(scene *world.Scene){

	// for i, c := range *scene.Map{
		// fmt.Println( *scene.Map )
	// }

}