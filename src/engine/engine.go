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

var resPath string

// Init engine
func Init(info def.LoadInfo){
	fmt.Println("Engine init...")
	
	resPath = info.ResourceFolder

	ui.Init(info.ScreenSize)

	cells, tsxName := loaders.LoadTmx(resPath+info.MapName)
	tileFileName, tileW, tileH := loaders.LoadTSX( resPath+tsxName)	
	
	entity.SetHero( &hero )
	entity.SetMap( cells )
	ui.LoadTiles(resPath+tileFileName, tileW, tileH)

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
