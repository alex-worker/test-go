package SDLRenderSystem

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	. "test-go/src/math"
)

func InitSDL(size Size2D) *sdl.Renderer {
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
