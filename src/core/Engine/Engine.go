package Engine

import (
	"fmt"
	. "test-go/src/core/FileManager"
	. "test-go/src/core/TileMap/Map"
	. "test-go/src/core/TileMap/TileAnimations"
	"test-go/src/core/TileMap/parser"
	. "test-go/src/core/sdl/SDLInputSystem"
	. "test-go/src/core/sdl/SDLRenderSystem"
	"test-go/src/defines"
	. "test-go/src/interfaces/IEngine"
	. "test-go/src/interfaces/IInputSystem"
	. "test-go/src/interfaces/IRenderSystem"
	. "test-go/src/interfaces/IResourceSystem"
)

type Engine struct {
	renderSystem   IRenderSystem
	resourceSystem IResourceSystem
	inputSystem    IInputSystem
}

func (e *Engine) Run() {
	fmt.Println("Engine::Run...")
	for {
		e.renderSystem.Draw()
		evt := e.inputSystem.GetInput()
		if evt == defines.EventQuit {
			break
		}
	}
}

func GetEngine(dataPath string) IEngine {
	resourceSystem := GetFileManager(dataPath)

	windowSize := defines.Size{Width: 640, Height: 480}

	renderSystem, err := GetRenderSystem(windowSize)
	if err != nil {
		panic(err)
	}

	inputSystem, err := GetInputSystem()
	if err != nil {
		panic(err)
	}

	eng := &Engine{
		resourceSystem: resourceSystem,
		renderSystem:   renderSystem,
		inputSystem:    inputSystem,
	}

	mapName := "swamp.tmx"

	tmxBuf, err := GetFile(&resourceSystem, mapName)
	if err != nil {
		panic(err)
	}

	tmx, err := parser.Parse(tmxBuf)
	if err != nil {
		panic(err)
	}

	_, err = LoadMap(tmx)
	if err != nil {
		panic(err)
	}

	animInfo := LoadTileSets(tmx)

	if len(animInfo) > 1 {
		panic("TileSets more then one not supported")
	}

	return eng
}
