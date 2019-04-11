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
			entity.HeroMove( def.DirDown )
		case def.EventPressUp:
			entity.HeroMove( def.DirUp )
		case def.EventPressLeft:
			entity.HeroMove( def.DirLeft )
		case def.EventPressRight:
			entity.HeroMove( def.DirRight )
		}

		entity.HeroMove( def.DirRight)

		ui.LookAtHero( entity.GetMap(), entity.GetHero() )

	}

	ui.Destroy()

}
