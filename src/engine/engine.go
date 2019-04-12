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

var cells *[][]def.Cell
var tiles *[]def.Tile

// Init engine
func Init(info def.LoadInfo){
	fmt.Println("Engine init...")

	def.SetResourceFolder( info.ResourceFolder )

	ui.Init(info.ScreenSize)

	var (
		tileFileName string
		tileW int32
		tileH int32
	)

	cells, tiles, tileFileName, tileW, tileH = loaders.LoadTmx(info.MapName)
	// tileFileName, tileW, tileH := loaders.LoadTSX( resPath+tsxName)

	entity.SetHero( &hero )
	entity.SetMap( cells )

	ui.LoadFont(info.FontName)
	
	if ( tileFileName != "" ){
		ui.LoadTiles(tileFileName, tileW, tileH)
	} else {
		println( "tileFileName is empty!" )
	}

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

		ui.LookAtHero( entity.GetMap(), entity.GetHero() )

	}

	ui.Destroy()

}
