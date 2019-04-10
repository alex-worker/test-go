package entity

import (
	"fmt"
)

type Cell uint32

var cells *[][]Cell
var mapX uint32 // размер карты
var mapY uint32

func LoadMap(filename string){
	fmt.Println("Load map", filename)
	cells, mapX, mapY = loadTmxMap( filename )
	fmt.Println( mapX, mapY, cells )
}

func GetMap() *[][]Cell{
	return cells
}