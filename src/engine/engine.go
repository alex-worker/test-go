package engine

import (
	"fmt"
	"./entity"
	"./ui"
	"./def"
	"./loaders"
)

var resPath string

// Init engine
func Init(info def.LoadInfo){
	fmt.Println("Engine init...")
	
	resPath = info.ResourceFolder

	ui.Init(info.ScreenSize)

	cells, tsxName := loaders.LoadTmx(resPath+info.MapName)
	tileFileName, tileW, tileH := loaders.LoadTSX( resPath+tsxName)

	println( tileFileName, tileW, tileH )
	entity.SetMap( cells )

	// ui.LoadTiles(info.TileName)
	// entity.LoadMap(info.MapName)
}

// Run цикл
func Run(){

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
