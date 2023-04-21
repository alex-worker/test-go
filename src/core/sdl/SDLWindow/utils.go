package SDLWindow

import (
	"bytes"
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"image/png"
	"test-go/src/defines"
	"unsafe"
)

func initSDL(size defines.Size) *sdl.Renderer {
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

	err = img.Init(img.INIT_PNG)
	if err != nil {
		panic(err)
	}

	err = ttf.Init()
	if err != nil {
		panic(err)
	}

	scrPixelWidth := size.Width
	scrPixelHeight := size.Height

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(scrPixelWidth), int32(scrPixelHeight),
		sdl.WINDOW_SHOWN,
		//sdl.WINDOW_OPENGL,
	)
	if err != nil {
		panic(err)
	}

	// SDL_SetHint(SDL_HINT_RENDER_VSYNC, "1" or "0");
	// sdl.SetHint(sdl.HINT_RENDER_VSYNC, "1")

	// SDL_RENDERER_ACCELERATED для хардварной поддержки
	// renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE )
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(err)
	}

	return renderer
	//backScreen, err = renderer.CreateTexture(sdl.PIXELFORMAT_BGR888, sdl.TEXTUREACCESS_TARGET, int32(scrPixelWidth), int32(scrPixelHeight))
	//if err != nil {
	//	panic(err)
	//}

	// sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")
}

func drawStart(renderer *sdl.Renderer) {
	err := renderer.Clear()
	if err != nil {
		panic(err)
	}
}

func drawEnd(renderer *sdl.Renderer) {

	//if isShowFps {
	//	showFPS()
	//}
	//
	renderer.Present()

	//currentTime := sdl.GetTicks64()
	//deltaTime := currentTime - startTicks

	//fmt.Println(deltaTime)

	//var fps uint64

	//if deltaTime > 0 {
	//	fps = 1000 / deltaTime
	//} else {
	//	fps = 0
	//}
	//
	//fmt.Println("fps: ", fps)
	//
	//lastTime = currentTime
	//
	//UpdateUI(deltaTime)
}

func PngFileToTexture(renderer *sdl.Renderer, buf *[]byte) (texture *sdl.Texture, size defines.Size) {
	bufReader := bytes.NewReader(*buf)

	myImage, err := png.Decode(bufReader)
	if err != nil {
		panic(err)
	}

	w := myImage.Bounds().Max.X
	h := myImage.Bounds().Max.Y

	pixels := make([]byte, w*h*4)
	bIndex := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, a := myImage.At(x, y).RGBA()
			pixels[bIndex] = byte(r / 256)
			bIndex++
			pixels[bIndex] = byte(g / 256)
			bIndex++
			pixels[bIndex] = byte(b / 256)
			bIndex++
			pixels[bIndex] = byte(a / 256)
			bIndex++
		}
	}

	texture, err = renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, int32(w), int32(h))
	if err != nil {
		panic(err)
	}

	err = texture.Update(nil, unsafe.Pointer(&pixels[0]), w*4)
	if err != nil {
		panic(err)
	}

	err = texture.SetBlendMode(sdl.BLENDMODE_BLEND)
	if err != nil {
		panic(err)
	}

	return texture, defines.Size{Width: uint32(w), Height: uint32(h)}
}
