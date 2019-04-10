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
	entity.LoadMap( info.MapName)
	fmt.Println(info)

}

func Run(){
	
	for true {
		if !ui.Update() {
			break
		}
	}
	
}