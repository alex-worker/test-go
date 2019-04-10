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
	fmt.Println(info)
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
			entity.HeroRotMove( def.DirDown )
		case def.EventPressUp:
			entity.HeroRotMove( def.DirUp )
		case def.EventPressLeft:
			entity.HeroRotMove( def.DirLeft )
		case def.EventPressRight:
			entity.HeroRotMove( def.DirRight )
		}

		ui.LookAtHero( entity.GetMap(), entity.GetHero() )

	}

	ui.Destroy()

}
