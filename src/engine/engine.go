package engine

import (
	defines "test-go/src/engine/defines"
	"test-go/src/engine/entity"
	"test-go/src/engine/loaders"
	log "test-go/src/engine/logger"
	"test-go/src/engine/resource"
	"test-go/src/engine/ui"
)

type Engine struct {
	hero defines.Hero
}

var hero = defines.Hero{
	Pos: defines.Pos{X: 0, Y: 0},
	Dir: defines.DirDown}

// размеры экрана в тайлах
var scrTilesSize = defines.Size{
	Width: 15, Height: 11,
}

var needReview bool

// var tiles *[]def.Tile

var logger log.ConsoleLogger

var view ui.View
var myMap defines.Map

// Init engine
func Create(info defines.LoadInfo) *Engine {
	logger.Log("Engine init...")
	engine := new(Engine)

	resource.SetResourceFolder(info.ResourceFolder)

	ui.Init(info.ScreenSize)

	var tilesets *[]loaders.TileSetInfo
	myMapPtr, tilesets := loaders.LoadTmx(info.MapName)

	myMap = *myMapPtr

	// fmt.Println( myMap.Layers[1].Data )

	loadTiles(tilesets)

	entity.SetHero(&hero)
	entity.SetMap(&myMap)

	ui.LoadFont(info.FontName)

	view.MakeView(&myMap, hero.Pos, scrTilesSize)
	return engine
}

func updateGame() bool {

	evt := ui.GetInput()

	if evt == defines.EventQuit {
		return false
	}

	switch evt {
	case defines.EventPressDown:
		entity.HeroDo(defines.DirDown, defines.ActionStand)
		needReview = true
	case defines.EventPressUp:
		entity.HeroDo(defines.DirUp, defines.ActionStand)
		needReview = true
	case defines.EventPressLeft:
		entity.HeroDo(defines.DirLeft, defines.ActionStand)
		needReview = true
	case defines.EventPressRight:
		entity.HeroDo(defines.DirRight, defines.ActionStand)
		needReview = true
	}

	if needReview {
		view.MakeView(&myMap, hero.Pos, scrTilesSize)
		needReview = false
	}
	return true
}

func drawGame() {
	ui.DrawStart()
	ui.DrawView(&view)
	ui.DrawEnd(true)
}

// RunOnce цикл
func (e Engine) RunOnce() {
	logger.Log("Engine run once...")
	updateGame()
	drawGame()
	ui.Destroy()
	logger.Log("Have a nice day!..")
}

// Run цикл
func (e Engine) Run() {

	logger.Log("Engine run...")

	for true {

		// start := ui.GetTickCount()
		if !updateGame() {
			break
		}
		drawGame()

	}

	ui.Destroy()
	logger.Log("Have a nice day!..")

}

func loadTiles(tilesets *[]loaders.TileSetInfo) {

	// TODO на будущее как-то надо подправить если текстур вдруг будет несколько
	for _, tileset := range *tilesets {
		ui.LoadTileset(
			tileset.Filename,
			tileset.TileW,
			tileset.TileH,
			tileset.Tiles,
		)
	}

}
