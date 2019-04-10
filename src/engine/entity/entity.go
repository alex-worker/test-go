package entity

import (
	"fmt"
)
type Cell uint32

// type Scene struct {
// 	Map *[][]Cell
// 	Width uint32
// 	Height uint32
// }

func LoadMap(filename string){
	fmt.Println("Load map", filename)
	mymap, w, h := loadTmxMap( filename )
	fmt.Println( w, h )
	fmt.Println( mymap )
}