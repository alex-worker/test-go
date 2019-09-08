package engine

import (
	"./def"
	"./entity"
	"./loaders"
	"./ui"
	"fmt"
)

var hero = def.Hero{
	Pos: def.Pos{X: 0, Y: 0},
	Dir: def.DirDown}

// размеры экрана в тайлах
var scrTilesSize = def.Size{
	Width: 15, Height: 11,
}

var needReview bool

// var tiles *[]def.Tile

var view ui.View
var myMap def.Map

// Init engine
func Init(info def.LoadInfo) {
	fmt.Println("Engine init...")

	def.SetResourceFolder(info.ResourceFolder)

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

}

func updateGame() bool {

	evt := ui.GetInput()

	if evt == def.EventQuit {
		return false
	}

	switch evt {
	case def.EventPressDown:
		entity.HeroDo(def.DirDown, def.ActionStand)
		needReview = true
	case def.EventPressUp:
		entity.HeroDo(def.DirUp, def.ActionStand)
		needReview = true
	case def.EventPressLeft:
		entity.HeroDo(def.DirLeft, def.ActionStand)
		needReview = true
	case def.EventPressRight:
		entity.HeroDo(def.DirRight, def.ActionStand)
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
func RunOnce() {
	fmt.Println("Engine run once...")
	updateGame()
	drawGame()
	ui.Destroy()
	fmt.Println("Have a nice day!..")
}

// Run цикл
func Run() {

	fmt.Println("Engine run...")

	for true {

		// start := ui.GetTickCount()
		if !updateGame() {
			break
		}
		drawGame()

	}

	ui.Destroy()
	fmt.Println("Have a nice day!..")

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
