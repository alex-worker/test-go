package ui

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	// "github.com/veandco/go-sdl2/mix"
	"../def"
)

var fps uint32
var deltaTime uint32
var lastTime uint32

var tileW int32
var tileH int32
var tileShift uint

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
var curFont *ttf.Font

// сдвиг на карте когда центрируемся на герое
var mapPosX int
var mapPosY int

// Destroy уничтожаем ui
func Destroy() {
	
	curFont.Close()
	ttf.Quit()

	renderer.Destroy()
	window.Destroy()
	sdl.Quit()
	fmt.Println("Ui offline...")
}

// Init инициализируем ui
func Init(scr def.Rect) {
	fmt.Println("UI Init...")
	// sdl.LogSetAllPriority(sdl.LOG_PRIORITY_VERBOSE)
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	img.Init(img.INIT_PNG)

	// err = mix.Init(mix.INIT_MP3)
	// if err != nil {
		// panic(err)
	// }

	err = ttf.Init()
	if err != nil {
		panic(err)
	}

	scrPixelWidth = scr.Width
	scrPixelHeight = scr.Height

	tilePixelSize = scrPixelWidth / scrTilesWidth

	// fmt.Println( tilePixelSize )

	window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(scrPixelWidth), int32(scrPixelHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	
	// SDL_SetHint(SDL_HINT_RENDER_VSYNC, "1" or "0");
	// sdl.SetHint(sdl.HINT_RENDER_VSYNC, "1")


	// SDL_RENDERER_ACCELERATED для хардварной поддержки
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE )
	if err != nil {
		panic(err)
	}

	// sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	keyboardState = sdl.GetKeyboardState()
}

// LoadFont грузим шрифт
func LoadFont(fontname string) {
	fmt.Println( "Loading font...", fontname)
	font,err := ttf.OpenFont(def.GetPath(fontname+".ttf"), 24)
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, "OpenFont: %s\n", err)
	}
	curFont = font
}

// LoadTiles загрузить файл тайлов
func LoadTiles(filename string, w int32, h int32) {
	tileW = w
	tileH = h

	texture, texW, _ := imgFileToTexture(filename)

	tileInTexture := texW / int(tileW)
	tileShift = 1

	for tileInTexture > 2 { // ручной логарифм по основанию 2 !
		tileInTexture = tileInTexture / 2
		tileShift++
	}

	textureAtlas = texture
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
func DrawTile(cell def.Cell, x int, y int) {

	mapY := int32(cell) >> tileShift
	mapX := int32(cell) - mapY<<tileShift

	srcRect := sdl.Rect{X: int32(mapX * tileW), Y: int32(mapY * tileH), W: tileW, H: tileH}
	dstRect := sdl.Rect{X: int32(x * tilePixelSize), Y: int32(y * tilePixelSize), W: int32(tilePixelSize), H: int32(tilePixelSize)}

	renderer.Copy(textureAtlas, &srcRect, &dstRect)

}

// LookAtHero рисуем карту и героя
func LookAtHero(cells *[][]def.Cell, hero *def.Hero) {

	mymap := *cells

	mapWidth := len(mymap)
	mapHeight := len(mymap[0])

	// половина экрана в тайлах
	scrHalfWidth := scrTilesWidth / 2
	scrHalfHeight := scrTilesHeight / 2

	// максимальное смещение
	scrWindowPosMaxX := mapWidth - scrTilesWidth
	scrWindowPosMaxY := mapHeight - scrTilesHeight

	mapPosX = hero.Pos.X - scrHalfWidth
	mapPosY = hero.Pos.Y - scrHalfHeight

	if mapPosX < 0 {
		mapPosX = 0
	}

	if mapPosY < 0 {
		mapPosY = 0
	}

	if mapPosX > scrWindowPosMaxX {
		mapPosX = scrWindowPosMaxX
	}

	if mapPosY > scrWindowPosMaxY {
		mapPosY = scrWindowPosMaxY
	}

	// println( mapPosX, mapPosY)

	renderer.Clear()

	for x := 0; x < scrTilesWidth; x++ {
		for y := 0; y < scrTilesHeight; y++ {
			cell := mymap[y+mapPosY][x+mapPosX]
			if cell != 0 {
				DrawTile(cell, x, y)
			}
		}
	}

	currentTime := sdl.GetTicks()

	deltaTime = currentTime - lastTime

	if deltaTime > 0 {
		fps = 1000 / deltaTime
	} else {
		fps = 0
	}

	fpsStr := fmt.Sprintf("fps: %v %v", fps, deltaTime) 
	// println(fps)
	printAt( fpsStr, 0, 0)

	lastTime = currentTime

	renderer.Present()
}


func printAt(text string, x int32, y int32){
	
	grapText, err := renderText( text,
		sdl.Color{R:255, G:255, B:255, A:0},
	)
	if err != nil {
		panic(err)
	}

	_, _, tW, tH, _ := grapText.Query()

	// const n = 128

	rect := sdl.Rect{ X: 10, Y:10, W: tW, H: tH}

	grapText.SetAlphaMod(255)
	renderer.Copy(grapText, nil, &rect)

}

func renderText(text string, color sdl.Color) (texture *sdl.Texture, err error) {
	surface, err := curFont.RenderUTF8Blended(text, color)
	if err != nil {
		panic(err)
	}

	defer surface.Free()

	texture, err = renderer.CreateTextureFromSurface(surface)
	return
}