package engine

import (
	"fmt"
	"./entity"
	"./ui"
	"./def"
)

// Init engine
func Init(info def.LoadInfo){
	fmt.Println("Engine init...")
	ui.Init(info.ScreenSize)
	ui.LoadTiles(info.TileName)
	entity.LoadMap(info.MapName)
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
