package ui

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Ui struct {
	Window *sdl.Window
	Surface *sdl.Surface
}

func ( *Ui ) Update() bool {

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		default:
			// evt := event.(type)
			fmt.Println( event.GetTimestamp() )
			break
		case *sdl.QuitEvent:
			println("Quit")
			// running = false
			return false
			break
		}
	}

	return true
}

func ( ui *Ui) Destroy(){
	ui.Window.Destroy()
	sdl.Quit()
	fmt.Println("Ui offline...")
}

func ( ui *Ui) Init(){

	// sdl.LogSetAllPriority(sdl.LOG_PRIORITY_VERBOSE)

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	// defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
    800, 600, sdl.WINDOW_SHOWN )
    if err != nil {
        panic(err)
    }
	// defer window.Destroy()

	ui.Window = window
	
	surface, err := window.GetSurface()
    if err != nil {
        panic(err)
    }
	ui.Surface = surface

	ui.Surface.FillRect(nil, 0)
	// surface.FillRect(nil, 0)

    // rect := sdl.Rect{0, 0, 200, 200}
    // surface.FillRect(&rect, 0xffff0000)
    window.UpdateSurface()

}