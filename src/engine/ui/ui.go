package ui

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"../def"
)

// размеры экрана в пикселях
var scrPixelWidth int
var scrPixelHeight int

// размеры экрана в тайлах
var scrTilesWidth = 15
var scrTilesHeight = 11

// размер одного тайла в пикселях
var tilePixelSize int

var window *sdl.Window
var renderer *sdl.Renderer
var textureAtlas *sdl.Texture

// сдвиг на карте когда центрируемся на герое
var mapPosX int
var mapPosY int

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
	
		scrPixelWidth = scr.Width
		scrPixelHeight = scr.Height

		tilePixelSize = scrPixelWidth / scrTilesWidth

		// fmt.Println( tilePixelSize )

		window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(scrPixelWidth), int32(scrPixelHeight), sdl.WINDOW_SHOWN )
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
func DrawTile(cell def.Cell, x int, y int){

	mapY := int(cell) >> 4
	mapX := int(cell) - mapY << 4

	srcRect := sdl.Rect{ X:int32(mapX*16), Y:int32(mapY*16), W:16, H:16 }
	dstRect := sdl.Rect{ X:int32(x*tilePixelSize), Y:int32(y*tilePixelSize), W:int32(tilePixelSize), H:int32(tilePixelSize) }

	renderer.Copy( textureAtlas, &srcRect, &dstRect )

}

// LookAtHero рисуем карту и героя
func LookAtHero(cells *[][]def.Cell, hero *def.Hero){

	// mapWidth := len( *cells )
	// mapWidthHalf := mapWidth / 32

// половина экрана в тайлах
	scrHalfWidth := scrTilesWidth / 2

// максимальное смещение 
	scrWindowPosMax := scrTilesWidth - scrHalfWidth

	mapPosX = hero.X - scrHalfWidth
	mapPosY = hero.Y - scrHalfWidth

	if ( mapPosX < 0 ) {
		mapPosX = 0
	}

	if ( mapPosY < 0 ) {
		mapPosY = 0
	}

	if ( mapPosX > scrWindowPosMax ){
		mapPosX = scrWindowPosMax
	}

	if ( mapPosY > scrWindowPosMax ){
		mapPosY = scrWindowPosMax
	}

	renderer.Clear()
	for x := 0; x < scrTilesWidth; x++ {
		for y := 0 ; y < scrTilesHeight; y++ {
			cell := (*cells)[y+mapPosX][x+mapPosY]
			DrawTile( cell, x, y)
		}
	}

	renderer.Present()
}
