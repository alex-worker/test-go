package entity

import (
	"fmt"
	"../def"
)

var cells *[][]def.Cell
var mapX uint32 // размер карты
var mapY uint32

func LoadMap(filename string){
	fmt.Println("Load map", filename)
	cells, mapX, mapY = loadTmxMap( filename )
	// fmt.Println( mapX, mapY, cells )
}

// а может в common все перекинуть? %)

func GetMap() *[][]def.Cell{
	return cells
}