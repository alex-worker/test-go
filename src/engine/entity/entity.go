package entity

import (
	"fmt"
	"../def"
)

var cells *[][]def.Cell
var mapX uint32 // размер карты
var mapY uint32

// LoadMap загрузить карту из файла
func LoadMap(filename string){
	fmt.Println("Load map", filename)
	cells, mapX, mapY = loadTmxMap( filename )
	// fmt.Println( mapX, mapY, cells )
}

// GetMap получаем карту
func GetMap() *[][]def.Cell{
	return cells
}