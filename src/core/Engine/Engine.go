package Engine

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	. "test-go/src/core"
	. "test-go/src/core/FileManager"
	. "test-go/src/core/sdl/SDLWindow"
	"test-go/src/defines"
	. "test-go/src/interfaces/IEngine"
	. "test-go/src/interfaces/IResourceManager"
	. "test-go/src/interfaces/IWindow"
)

type Engine struct {
	resourceManager IResourceManager
	window          IWindow
	fps             uint64
}

func (e *Engine) Run() {
	fmt.Println("Engine::Run...")
	for {
		startTicks := sdl.GetTicks64()
		e.window.Update()
		endTicks := sdl.GetTicks64()
		e.fps = CalcFPS(startTicks, endTicks)
		println(e.fps)
		evt := e.window.GetInput()
		if evt == defines.EventQuit {
			break
		}
	}
}

func GetEngine() IEngine {
	resourceManager := GetFileManager("./data")

	windowSize := defines.Size{Width: 640, Height: 480}

	win, err := GetWindow(windowSize)
	if err != nil {
		panic(err)
	}

	eng := &Engine{
		resourceManager: resourceManager,
		window:          win,
	}

	return eng
}

func getFile(r *IResourceManager, name string) *[]byte {
	res, err := (*r).GetResource(name)
	if err != nil {
		panic(err)
	}
	buf, err := res.GetContent()
	if err != nil {
		panic(err)
	}
	return buf
}
