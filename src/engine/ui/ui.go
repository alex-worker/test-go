package ui

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

var window *sdl.Window = nil
var renderer *sdl.Renderer = nil
var textureAtlas *sdl.Texture = nil

func Destroy(){
	renderer.Destroy()
	window.Destroy()
	sdl.Quit()
	fmt.Println("Ui offline...")
}

func Init(){
	fmt.Println("UI Init...")
		// sdl.LogSetAllPriority(sdl.LOG_PRIORITY_VERBOSE)
		err := sdl.Init(sdl.INIT_EVERYTHING)
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
	
		// sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")
		// textureAtlas = imgFileToTexture("ui/assets/tiles_many.png")
}

func LoadTiles(filename string){

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