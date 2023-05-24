package Engine

import (
	"fmt"
	. "test-go/src/core/FileManager"
	. "test-go/src/core/TileMap/Map"
	. "test-go/src/core/TileMap/TileAnimations"
	"test-go/src/core/TileMap/parser"
	. "test-go/src/core/sdl/SDLInputSystem"
	. "test-go/src/core/sdl/SDLRenderSystem"
	"test-go/src/core/sdl/SDLViewMap2D"
	"test-go/src/defines"
	. "test-go/src/math"
)

type Engine struct {
	renderSystem   *SDLRenderSystem
	resourceSystem *FileManager
	inputSystem    *SDLInputSystem
	mapView        *SDLViewMap2D.SDLViewMap2D
}

func (e *Engine) Run() error {
	fmt.Println("Engine::Run...")
	for {
		e.renderSystem.Draw()
		err := e.mapView.Draw(e.renderSystem)
		if err != nil {
			return err
		}
		evt := e.inputSystem.GetInput()
		if evt == defines.EventQuit {
			break
		}
	}
	return nil
}

func GetEngine(dataPath string) (*Engine, error) {
	resourceSystem, err := GetFileManager(dataPath)
	if err != nil {
		panic(err)
	}

	windowSize := Size2D{Width: 640, Height: 480}

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

	eng.mapView, err = eng.parseMap(mapName)
	if err != nil {
		panic(err)
	}

	return eng, nil
}

func (e *Engine) parseMap(mapName string) (*SDLViewMap2D.SDLViewMap2D, error) {
	tmxBuf, err := GetFile(e.resourceSystem, mapName)
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

	textureBuf, err := GetFile(e.resourceSystem, animInfo[0].FileName)
	if err != nil {
		return nil, err
	}

	texture, err := e.renderSystem.GetTexture(textureBuf)
	if err != nil {
		return nil, err
	}

	viewSize := Size2D{
		Width:  10,
		Height: 10,
	}
	return SDLViewMap2D.New(viewSize, texture.Texture)
}
