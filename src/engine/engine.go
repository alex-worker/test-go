package engine

import (
	"fmt"
	"./entity"
	"./ui"
)

type LoadInfo struct {
	MapName string
	TileName string
}

func Init(info LoadInfo){
	fmt.Println("Engine init...")
	ui.Init()
	ui.LoadTiles(info.TileName)
	entity.LoadMap(info.MapName)
	fmt.Println(info)
}

func Run(){
	
	for true {
		ui.Draw( entity.GetMap() )
		if !ui.Update() {
			break
		}
	}

	ui.Destroy()
	
}