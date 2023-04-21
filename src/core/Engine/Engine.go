package Engine

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	. "test-go/src/core"
	. "test-go/src/core/FileManager"
	"test-go/src/core/TileMap"
	"test-go/src/core/TileMap/parser"
	. "test-go/src/core/sdl/SDLWindow"
	"test-go/src/defines"
	. "test-go/src/interfaces/IEngine"
	. "test-go/src/interfaces/IResourceManager"
)

type Engine struct {
	resourceManager IResourceManager
	window          *SDLWindow
	fps             uint64
}

func (e *Engine) Run() {
	fmt.Println("Engine::Run...")
	for {
		startTicks := sdl.GetTicks64()
		e.window.DrawStart()
		e.window.DrawEnd()
		endTicks := sdl.GetTicks64()
		e.fps = CalcFPS(startTicks, endTicks)
		//println(e.fps)
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

	tmxBuf, err := GetFile(&resourceManager, "mycastle.tmx")
	if err != nil {
		panic(err)
	}

	tmx, err := parser.Parse(tmxBuf)
	if err != nil {
		panic(err)
	}

	m, err := TileMap.Load(tmx)
	if err != nil {
		panic(err)
	}

	fmt.Println(m)

	return eng
}
