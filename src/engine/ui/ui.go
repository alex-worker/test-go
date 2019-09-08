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
var scrPixelWidth uint32
var scrPixelHeight uint32

// размеры экрана в тайлах
var scrTilesWidth = uint32(15)
var scrTilesHeight = uint32(11)

// размер одного тайла в пикселях
var tilePixelSize uint32

var window *sdl.Window
var renderer *sdl.Renderer

// var surface *sdl.Surface
var backScreen *sdl.Texture
var textureAtlas *sdl.Texture
var curFont *ttf.Font

var animateTiles *def.AnimateTiles

var view View

// Destroy уничтожаем ui
func Destroy() {

	curFont.Close()
	ttf.Quit()

	err := renderer.Destroy()
	if err != nil {
		panic(err)
	}

	err = window.Destroy()
	if err != nil {
		panic(err)
	}

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
		_, err = sdl.GetRenderDriverInfo(i, &drinfo)
		if err != nil {
			panic(err)
		}
		println("Driver name", drinfo.Name)
	}

	img.Init(img.INIT_PNG)

	err = ttf.Init()
	if err != nil {
		panic(err)
	}

	scrPixelWidth = scr.Width
	scrPixelHeight = scr.Height

	tilePixelSize = scrPixelWidth / scrTilesWidth

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

// LoadTileset загрузить текстуру и запомнить анимацию
// пока так, потом посмотрим
func LoadTileset(filename string, w int32, h int32, anim *def.AnimateTiles) {
	animateTiles = anim
	textureAtlas = loadTexture(filename, w, h)
}

// LoadTexture загрузить файл тайлов
func loadTexture(filename string, w int32, h int32) *sdl.Texture {
	tileW = w
	tileH = h

	texture, texW, _ := imgFileToTexture(filename)

	tileInTexture := texW / int(tileW)
	tileShift = 1

	for tileInTexture > 2 { // ручной логарифм по основанию 2 !
		tileInTexture = tileInTexture / 2
		tileShift++
	}

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
func DrawStart() {
	err := renderer.Clear()
	if err != nil {
		panic(err)
	}
}

// DrawEnd Окончить отрисовку
func DrawEnd(isShowFps bool) {

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

	UpdateUI(deltaTime)

}

// timing - отрицательное число которое осталось от предыдущего тайминга
func nextTile(timing int, t *def.AnimateTile) {
	t.Index++
	if t.Index > len(t.Frames)-1 {
		t.Index = 0
	}
	t.Tick = int(t.Frames[t.Index].Duration) + timing
}

func updateTile(delta uint32, t *def.AnimateTile) {
	t.NeedUpdate = false
	timing := t.Tick - int(delta)

	if timing < 0 {
		nextTile(timing, t)
	} else {
		t.Tick = timing
	}
}

func updateAnimation(delta uint32) {

	// var animateTiles *def.AnimateTiles
	for _, t := range *animateTiles {
		if t.NeedUpdate {
			updateTile(delta, t)
			// println("anim!")
		}
	}
}

// ShowFPS показать FPS
func showFPS() {

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

	err = grapText.SetAlphaMod(255)
	if err != nil {
		panic(err)
	}

	err = renderer.Copy(grapText, nil, &rect)
	if err != nil {
		panic(err)
	}

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

// GetTicks время со старта игры
func GetTicks() uint32 {
	return sdl.GetTicks()
}

// Delay спим!
func Delay(time uint32) {
	println("Delay", time)
	sdl.Delay(time)
}

// DrawView рисуем окошко героя
func DrawView(v *View) {

	for _, layer := range v.Layers {
		drawLayer(layer, v.Size)
	}

}

func drawLayer(l *Layer, size def.Size) {

	layer := *l

	x := uint32(0)
	y := uint32(0)
	for index := 0; index < len(layer); index++ {
		drawTile(layer[index], def.Pos{X: x, Y: y})
		x++
		if x == size.Width {
			x = 0
			y++
		}
	}

}

func getAnimTile(c def.Cell, delta uint32) (tile def.Cell) {

	tile = c

	animate := *animateTiles

	if anim, ok := animate[c]; ok {
		anim.NeedUpdate = true
		index := anim.Index
		tile = anim.Frames[index].Cell
	}

	return
}

func drawTile(c def.Cell, pos def.Pos) {

	c = getAnimTile(c, deltaTime)

	mapY := int32(c) >> tileShift
	mapX := int32(c) - mapY<<tileShift

	srcRect := sdl.Rect{X: mapX * tileW, Y: mapY * tileH, W: tileW, H: tileH}
	dstRect := sdl.Rect{X: int32(pos.X * tilePixelSize), Y: int32(pos.Y * tilePixelSize), W: int32(tilePixelSize), H: int32(tilePixelSize)}
	err := renderer.Copy(textureAtlas, &srcRect, &dstRect)
	if err != nil {
		panic(err)
	}

}

// UpdateUI обновление UI
func UpdateUI(delta uint32) {
	updateAnimation(deltaTime)
}
