package ui

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"../def"
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

func Init(scr def.Rect){
	fmt.Println("UI Init...")
		// sdl.LogSetAllPriority(sdl.LOG_PRIORITY_VERBOSE)
		err := sdl.Init(sdl.INIT_EVERYTHING)
		if err != nil {
			panic(err)
		}
	
		window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(scr.Width), int32(scr.Height), sdl.WINDOW_SHOWN )
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
	textureAtlas = imgFileToTexture(filename)
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

func DrawTile(cell def.Cell, x uint32, y uint32){

	mapY := int(cell) >> 4
	mapX := int(cell) - mapY << 4

	srcRect := sdl.Rect{ int32(mapX*16), int32(mapY*16), 16, 16 }
	dstRect := sdl.Rect{ int32(x*32), int32(y*32), 32, 32 }

	renderer.Copy( textureAtlas, &srcRect, &dstRect )

}

func Draw(calls *[][]def.Cell){

	// w := scene.Width
	// h := scene.Height

	var w uint32 = 15
	var h uint32 = 11

	var x uint32 = 0
	var y uint32 = 0

	for x = 0 ; x < w; x++ {
		for y = 0 ; y < h; y++ {
			cell := (*calls)[y][x]
			DrawTile( cell, x, y)
		}
	}

	renderer.Present()
}