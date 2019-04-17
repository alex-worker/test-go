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
var tilePixelSize int32

var window *sdl.Window
var renderer *sdl.Renderer

// var surface *sdl.Surface
var backScreen *sdl.Texture
var textureAtlas *sdl.Texture
var curFont *ttf.Font

var view View

// сдвиг на карте когда центрируемся на герое
// var mapPosX int
// var mapPosY int

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
func Init(scr def.Size) {
	fmt.Println("UI Init...")
	// sdl.LogSetAllPriority(sdl.LOG_PRIORITY_VERBOSE)
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	numdrivers, _ := sdl.GetNumRenderDrivers()
	for i := 0; i < numdrivers; i++ {
		var drinfo sdl.RendererInfo
		sdl.GetRenderDriverInfo(i, &drinfo)
		println("Driver name", drinfo.Name)
	}

	img.Init(img.INIT_PNG)

	err = ttf.Init()
	if err != nil {
		panic(err)
	}

	scrPixelWidth = scr.Width
	scrPixelHeight = scr.Height

	tilePixelSize = int32(scrPixelWidth / scrTilesWidth)

	window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(scrPixelWidth), int32(scrPixelHeight),
		sdl.WINDOW_SHOWN,
		// sdl.WINDOW_OPENGL,
	)
	if err != nil {
		panic(err)
	}

	// SDL_SetHint(SDL_HINT_RENDER_VSYNC, "1" or "0");
	// sdl.SetHint(sdl.HINT_RENDER_VSYNC, "1")

	// SDL_RENDERER_ACCELERATED для хардварной поддержки
	// renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE )
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(err)
	}

	backScreen, err = renderer.CreateTexture(sdl.PIXELFORMAT_BGR888, sdl.TEXTUREACCESS_TARGET, int32(scrPixelWidth), int32(scrPixelHeight))
	if err != nil {
		panic(err)
	}

	// sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	keyboardState = sdl.GetKeyboardState()
}

// LoadFont грузим шрифт
func LoadFont(fontname string) {
	fmt.Println("Loading font...", fontname)
	font, err := ttf.OpenFont(def.GetPath(fontname+".ttf"), 24)
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, "OpenFont: %s\n", err)
	}
	curFont = font
}

// LoadTiles загрузить файл тайлов
func LoadTiles(filename string, w int32, h int32) *sdl.Texture {
	tileW = w
	tileH = h

	texture, texW, _ := imgFileToTexture(filename)

	tileInTexture := texW / int(tileW)
	tileShift = 1

	for tileInTexture > 2 { // ручной логарифм по основанию 2 !
		tileInTexture = tileInTexture / 2
		tileShift++
	}

	// texture.SetBlendMode( sdl.BLENDMODE_BLEND )
	textureAtlas = texture
	return texture
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

// DrawStart Начать отрисовку
func DrawStart(){
	renderer.Clear()
}

// DrawEnd Окончить отрисовку
func DrawEnd( isShowFps bool ){

	if isShowFps {
		showFPS()
	}

	renderer.Present()
	
	currentTime := sdl.GetTicks()

	deltaTime = currentTime - lastTime

	if deltaTime > 0 {
		fps = 1000 / deltaTime
	} else {
		fps = 0
	}

	lastTime = currentTime

}

// ShowFPS показать FPS
func showFPS(){

	fpsStr := fmt.Sprintf("fps: %v %v", fps, deltaTime)
	printAt(fpsStr, 0, 0)

}

func printAt(text string, x int32, y int32) {

	grapText, err := renderText(text,
		sdl.Color{R: 255, G: 255, B: 255, A: 0},
	)
	if err != nil {
		panic(err)
	}

	_, _, tW, tH, _ := grapText.Query()

	rect := sdl.Rect{X: 10, Y: 10, W: tW, H: tH}

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


/*

// DrawTile рисуем один тайл
func DrawTile(cell def.Cell, x int, y int) {

	// mapY := int32(cell) >> tileShift
	// mapX := int32(cell) - mapY<<tileShift

	// srcRect := sdl.Rect{X: mapX * tileW, Y: mapY * tileH, W: tileW, H: tileH}

	srcRect := sdl.Rect{X: 32, Y: 0, W: tileW, H: tileH}

	dstRect := sdl.Rect{X: tilePixelSize * int32(x), Y: tilePixelSize * int32(y), W: int32(tilePixelSize), H: int32(tilePixelSize)}

	renderer.Copy(textureAtlas, &srcRect, &dstRect)

}

func DirectDrawTile(cell int32, x int, y int) {

	// mapY := int32(cell) >> tileShift
	// mapX := int32(cell) - mapY<<tileShift

	// srcRect := sdl.Rect{X: mapX * tileW, Y: mapY * tileH, W: tileW, H: tileH}

	srcRect := sdl.Rect{X: cell, Y: 0, W: tileW, H: tileH}

	dstRect := sdl.Rect{X: tilePixelSize * int32(x), Y: tilePixelSize * int32(y), W: int32(tilePixelSize), H: int32(tilePixelSize)}

	renderer.Copy(textureAtlas, &srcRect, &dstRect)

}

func drawLayer(layer *def.Layer, mapPosX int, mapPosY int) {

	mydata := (*layer).Data
	mymap := *mydata

	// fmt.Println( mapPosX, mapPosY )
	for x := 0; x < scrTilesWidth; x++ {
		for y := 0; y < scrTilesHeight; y++ {
			cell := mymap[y+mapPosY][x+mapPosX]
			// if cell != 0 {
			// DrawTile(cell, x, y)
			// DirectDrawTile( int32(2), x, y)
			// }
			srcRect := sdl.Rect{X: int32(cell)+int32(x+mapPosX)*tileW, Y: int32(y+mapPosY)*tileH, W: tileW, H: tileH}
			dstRect := sdl.Rect{X: tilePixelSize * int32(x), Y: tilePixelSize * int32(y), W: int32(tilePixelSize), H: int32(tilePixelSize)}
			renderer.Copy(textureAtlas, &srcRect, &dstRect)
		
		}
	}

}

// возвращаем x и y левой верхней точки view карты так чтобы x и y было в центре и не
// выходил за границы карты
func getMapPos(mapWidth int, mapHeight int, pos def.Pos) (posX int, posY int) {
	// половина экрана в тайлах
	scrHalfWidth := scrTilesWidth / 2
	scrHalfHeight := scrTilesHeight / 2

	// максимальное смещение
	scrWindowPosMaxX := mapWidth - scrTilesWidth
	scrWindowPosMaxY := mapHeight - scrTilesHeight

	posX = pos.X - scrHalfWidth
	posY = pos.Y - scrHalfHeight

	// println( posX, posY )

	if posX < 0 {
		posX = 0
	}

	if posY < 0 {
		posY = 0
	}

	if posX > scrWindowPosMaxX {
		posX = scrWindowPosMaxX
	}

	if posY > scrWindowPosMaxY {
		posY = scrWindowPosMaxY
	}

	// println( posX, posY )
	return posX, posY
}

// GetTickCount время со старта игры
func GetTickCount() uint32 {
	return sdl.GetTicks()
}

// Delay спим!
func Delay(time uint32) {
	println("Delay", time)
	sdl.Delay(time)
}

// LookAtHero рисуем карту и героя
func LookAtHero(layers *map[string]*def.Layer, mapWidth int, mapHeight int, hero *def.Hero) {

	// mapPosX, mapPosY := getMapPos(mapWidth, mapHeight, hero.Pos)

	// renderer.Clear()
	// renderer.SetRenderTarget( backScreen )
	renderer.Clear()
	// sdl.Delay(10)
	// if err != nil {
	// 	panic(err)
	// }
	// renderer.Clear()

	// println( mapPosX, mapPosX )
	// for _, layer := range *layers {
		// drawLayer(layer, mapPosX, mapPosY)
	// }

	currentTime := sdl.GetTicks()

	deltaTime = currentTime - lastTime

	if deltaTime > 0 {
		fps = 1000 / deltaTime
	} else {
		fps = 0
	}

	fpsStr := fmt.Sprintf("fps: %v %v", fps, deltaTime)
	printAt(fpsStr, 0, 0)

	lastTime = currentTime

	// renderer.SetRenderTarget( nil )
	// renderer.Clear()
	// src :=  sdl.Rect{X: 100, Y: i, W: int32(scrPixelWidth), H: int32(scrPixelHeight) }
	// for i:=int32(0);i<100;i++ {
	// 	renderer.Clear()
	// 	dst := sdl.Rect{X: 100, Y: i, W: int32(scrPixelWidth), H: int32(scrPixelHeight) }
	// 	renderer.Copy(textureAtlas, &src, &dst)
	// 	renderer.Present()
	// }

	// rect := sdl.Rect{ X: 0, Y:0, W: int32(scrPixelWidth), H: int32(scrPixelHeight) }
	// renderer.CopyEx( )
	// backScreen.SetColorMod( 200,0,0 )
	// backScreen.SetBlendMode( sdl.BLENDMODE_NONE )
	// renderer.Copy( backScreen, nil, nil )
	renderer.Present()
	// sdl.GLSetSwapInterval(1)
	// window.GLSwap()

	// window.GetSurface()
	// window.UpdateSurface()
	// sdl.Delay(400)
	// renderer.Clear()
	// renderer.SetRenderTarget( backScreen )
}

*/

// DrawView рисуем окошко героя
func DrawView(v *View){

	// vSize := v.GetSize()

	// layer := v.Layers[0]

	for _, layer := range v.Layers {
		drawLayer( layer , v.Size )
	}

	// fmt.Println( v.W` )
	// fmt.Println( layer )

}

func drawLayer( l *Layer, size def.Size ){

	layer:= *l

	// index := 0
	x:=0
	y:= 0
	for index:=0; index< len(layer); index++ {
		drawTile( layer[index], def.Pos{ X:x, Y: y } )
		x++
		if x == size.Width {
			x = 0
			y++
		}
	}

}

func drawTile( c def.Cell, pos def.Pos ){
	mapY := int32(c) >> tileShift
	mapX := int32(c) - mapY<<tileShift

	srcRect := sdl.Rect{X: mapX * tileW, Y: mapY * tileH, W: tileW, H: tileH}

	// srcRect := sdl.Rect{X: 32, Y: 0, W: tileW, H: tileH}
	dstRect := sdl.Rect{X: int32(pos.X)*32, Y: int32(pos.Y)*32, W: tileW, H: tileH}

	// dstRect := sdl.Rect{X: tilePixelSize * int32(x), Y: tilePixelSize * int32(y), W: int32(tilePixelSize), H: int32(tilePixelSize)}

	renderer.Copy(textureAtlas, &srcRect, &dstRect)
}