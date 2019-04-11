package ui

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"../def"
)

var keyboardState []uint8

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

		keyboardState = sdl.GetKeyboardState()
}

// LoadTiles загрузить файл тайлов
func LoadTiles(filename string){
	textureAtlas = imgFileToTexture(filename)
}

// LoadTilesFromTiledMap загружаем по имени файла в карте 
func LoadTilesFromTiledMap(filename string){
	
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

	if keyboardState[sdl.SCANCODE_UP] != 0 {
		return def.EventPressUp
	} else if keyboardState[sdl.SCANCODE_DOWN] != 0 {
		return def.EventPressDown
	} else if keyboardState[sdl.SCANCODE_LEFT] != 0 {
		return def.EventPressLeft
	} else if keyboardState[sdl.SCANCODE_RIGHT] != 0 {
		return def.EventPressRight
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

	mymap := *cells

	mapWidth := len( mymap )
	mapHeight := len( mymap[0] ) 

	// println( mapHeight )

// половина экрана в тайлах
	scrHalfWidth := scrTilesWidth / 2
	scrHalfHeight := scrTilesHeight / 2 

// максимальное смещение 
	scrWindowPosMaxX := mapWidth - scrTilesWidth
	scrWindowPosMaxY := mapHeight - scrTilesHeight

	mapPosX = hero.Pos.X - scrHalfWidth
	mapPosY = hero.Pos.Y - scrHalfHeight

	if ( mapPosX < 0 ) {
		mapPosX = 0
	}

	if ( mapPosY < 0 ) {
		mapPosY = 0
	}

	if ( mapPosX > scrWindowPosMaxX ){
		mapPosX = scrWindowPosMaxX
	}

	if ( mapPosY > scrWindowPosMaxY ){
		mapPosY = scrWindowPosMaxY
	}

	// println( mapPosX, mapPosY)

	renderer.Clear()
	for x := 0; x < scrTilesWidth; x++ {
		for y := 0 ; y < scrTilesHeight; y++ {
			cell := mymap[y+mapPosY][x+mapPosX]
			DrawTile( cell, x, y)
		}
	}

	renderer.Present()
}
