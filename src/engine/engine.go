package engine

import (
	"fmt"
	"./entity"
	"./ui"
	"./def"
	"./loaders"
)

var hero = def.Hero{
	Pos: def.Pos{X:5, Y:5},
	Dir: def.DirDown }

// var resPath string

// var cells *[][]def.Cell
var layers *map[string]*def.Layer
var tiles *[]def.Tile

// Init engine
func Init(info def.LoadInfo){
	fmt.Println("Engine init...")

	def.SetResourceFolder( info.ResourceFolder )

	ui.Init(info.ScreenSize)

	// var (
	// 	tileFileName string
	// 	tileW int32
	// 	tileH int32
	// )

	var tilesets *map[string]loaders.TileSetInfo
	layers, tilesets = loaders.LoadTmx(info.MapName)

	loadTiles(tilesets)

	entity.SetHero(&hero)
	entity.SetMap(layers)

	ui.LoadFont(info.FontName)
	
	// if ( tileFileName != "" ){
	// 	ui.LoadTiles(tileFileName, tileW, tileH)
	// } else {
	// 	println( "tileFileName is empty!" )
	// }

}

// Run цикл
func Run(){

	fmt.Println("Engine run...")

	var evt def.GameEvent

	for true {

		evt = ui.GetInput()

		if evt == def.EventQuit {
			break
		}

		switch evt {
		case def.EventPressDown:
			entity.HeroDo( def.DirDown, def.ActionStand )
		case def.EventPressUp:
			entity.HeroDo( def.DirUp, def.ActionStand )
		case def.EventPressLeft:
			entity.HeroDo( def.DirLeft, def.ActionStand )
		case def.EventPressRight:
			entity.HeroDo( def.DirRight, def.ActionStand )
		}

		mymap, w, h := entity.GetMap()
		
		ui.LookAtHero( mymap, int(w), int(h), &hero )

	}

	ui.Destroy()

}

func loadTiles( tilesets *map[string]loaders.TileSetInfo ){

	for _, tileset := range *tilesets {
		ui.LoadTiles(tileset.Filename, tileset.TileW , tileset.TileH )
	}

}