package engine

import (
	"fmt"
	"./def"
	"./entity"
	"./loaders"
	"./ui"
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

// Run цикл
func Run() {

	fmt.Println("Engine run...")

	for true {

		// start := ui.GetTickCount()
		if !updateGame() {
			break
		}
		drawGame()

		// time := int(ui.GetTickCount()) - int(start);
		// // if (time < 0) continue; // if time is negative, the time probably overflew, so continue asap
		// if time>0 {
		// 	sleepTime := 1000/10 - time
		// 	if sleepTime > 0 {
		// 		// ui.Delay( uint32(sleepTime))
		// 	}
		// }
	}

	ui.Destroy()
	fmt.Println("Have a nice day!..")

}

func loadTiles(tilesets *[]loaders.TileSetInfo) {

// TODO на будущее как-то надо подправить если текстур вдруг будет несколько
	for _, tileset := range *tilesets {
		ui.LoadTexture(tileset.Filename, tileset.TileW, tileset.TileH)
	}

}
