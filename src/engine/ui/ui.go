package ui

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"../def"
)

// размеры экрана в пикселях
var scrPixelWidth int32
var scrPixelHeight int32

// размеры экрана в тайлах
var scrTilesWidth uint32 = 15
var scrTilesHeight uint32 = 11

// размер одного тайла в пикселях
var tilePixelSize uint32

var window *sdl.Window
var renderer *sdl.Renderer
var textureAtlas *sdl.Texture

// Destroy уничтожаем ui
func Destroy(){
	renderer.Destroy()
	window.Destroy()
	sdl.Quit()
	fmt.Println("Ui offline...")
}

// Init инициализируем ui
func Init(scr def.Rect){
	fmt.Println("UI Init...")
		// sdl.LogSetAllPriority(sdl.LOG_PRIORITY_VERBOSE)
		err := sdl.Init(sdl.INIT_EVERYTHING)
		if err != nil {
			panic(err)
		}
	
		scrPixelWidth = int32(scr.Width)
		scrPixelHeight = int32(scr.Height)

		tilePixelSize = uint32(scrPixelWidth) / scrTilesWidth

		// fmt.Println( tilePixelSize )

		window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		scrPixelWidth, scrPixelHeight, sdl.WINDOW_SHOWN )
		if err != nil {
			panic(err)
		}
		
		// SDL_RENDERER_ACCELERATED для хардварной поддержки
		renderer, err = sdl.CreateRenderer( window, -1, sdl.RENDERER_SOFTWARE)
		if err != nil {
			panic(err)
		}
	
		// sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")
}

// LoadTiles загрузить файл тайлов
func LoadTiles(filename string){
	textureAtlas = imgFileToTexture(filename)
}

// GetInput обновление событий экрана
// return false when window is closed
func GetInput() def.GameEvent {

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		default:
			// fmt.Println( event )
			break
		case *sdl.QuitEvent:
			return def.EventQuit
		}
	}

	return def.EventNo
}

// DrawTile рисуем один тайл
func DrawTile(cell def.Cell, x uint32, y uint32){

	mapY := int(cell) >> 4
	mapX := int(cell) - mapY << 4

	srcRect := sdl.Rect{ X:int32(mapX*16), Y:int32(mapY*16), W:16, H:16 }
	dstRect := sdl.Rect{ X:int32(x*tilePixelSize), Y:int32(y*tilePixelSize), W:int32(tilePixelSize), H:int32(tilePixelSize) }

	renderer.Copy( textureAtlas, &srcRect, &dstRect )

}

// LookAtHero рисуем карту и героя
func LookAtHero(calls *[][]def.Cell, hero *def.Hero){

	var x uint32
	var y uint32

	for x = 0 ; x < scrTilesWidth; x++ {
		for y = 0 ; y < scrTilesHeight; y++ {
			cell := (*calls)[y+hero.Y][x+hero.X]
			DrawTile( cell, x, y)
		}
	}

	renderer.Present()
}