package Engine

import (
	"fmt"
	. "test-go/src/core/FileManager"
	. "test-go/src/core/SDLWindow"
	"test-go/src/defines"
	. "test-go/src/interfaces/IEngine"
	. "test-go/src/interfaces/IResourceManager"
	. "test-go/src/interfaces/IWindow"
)

type Engine struct {
	resourceManager IResourceManager
	window          IWindow
}

func (e *Engine) Run() {
	fmt.Println("Engine::Run...")
	for {
		e.window.Update()
		evt := e.window.GetInput()
		if evt == defines.EventQuit {
			break
		}
	}
}

func GetEngine() IEngine {
	resourceManager := GetFileManager("/data")

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
