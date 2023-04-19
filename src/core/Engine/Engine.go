package Engine

import (
	. "test-go/src/core/FileManager"
	. "test-go/src/core/SDLWindow"
	. "test-go/src/interfaces/IEngine"
	. "test-go/src/interfaces/IResourceManager"
	. "test-go/src/interfaces/IWindow"
)

type Engine struct {
	resourceManager IResourceManager
	window          IWindow
}

func (e *Engine) Run() {
}

func GetEngine() IEngine {
	resourceManager := GetFileManager("/data")

	win, err := GetWindow()
	if err != nil {
		panic(err)
	}

	eng := &Engine{
		resourceManager: resourceManager,
		window:          win,
	}
	return eng
}
