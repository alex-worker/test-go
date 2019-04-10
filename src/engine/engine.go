package engine

import (
	"fmt"
	"./entity"
)

type LoadInfo struct {
	MapName string
	TileName string
}

func Init(info LoadInfo){
	fmt.Println("Engine init...")
	entity.LoadMap( info.MapName)
	fmt.Println(info)

}