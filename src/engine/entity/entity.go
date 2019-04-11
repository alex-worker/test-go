package entity

import (
	"fmt"
	"../def"
)

var cells *[][]def.Cell
var mapX uint32 // размер карты
var mapY uint32

var curHero *def.Hero

// SetMap устанавливаем карту
func SetMap( mycells *[][]def.Cell ){
	cells = mycells
	mapX = uint32(len (*mycells))
	mapY = uint32(len ((*mycells)[0]))
}

// SetHero устанавливаем главного героя
func SetHero( hero *def.Hero ){
	curHero = hero
}

// // LoadMap загрузить карту из файла
// func LoadMap(filename string){
// 	fmt.Println("Load map", filename)

// 	var tsxName string

// 	cells, mapX, mapY, tsxName = loadTmxMap( filename )
	
// 	hero := def.Hero{
// 		Pos: def.Pos{X:177, Y:542},
// 		Dir: def.DirDown }

// 	curHero = &hero
	
// 	// fmt.Println( mapX, mapY, cells )
// }

// GetMap получаем карту
func GetMap() *[][]def.Cell{
	return cells
}

// GetHero указатель на героя
func GetHero() *def.Hero {
	return curHero
}
 
// HeroDo герой что-то делает
func HeroDo( dir def.Direction, act def.HeroAction ){

	oldPos := curHero.Pos

	newPos,err := calcNewPos( &oldPos, dir, int(mapX), int(mapY) )
	if err != nil {
		fmt.Print(err)
		return
	}

	heroAction( newPos, def.ActionStand )

}

// // HeroMove герой поворачивается или двигается
// func HeroMove( dir def.Direction ){

// 	oldPos := curHero.Pos

// 	newPos,err := calcNewPos( &oldPos, dir )
// 	if err != nil {
// 		fmt.Print(err)
// 		return
// 	}

// 	// fmt.Println( newPos )

// 	heroAction( newPos, def.ActionStand )

// }

func heroAction(pos *def.Pos, act def.HeroAction){
	if( act == def.ActionStand ){
		heroStand(pos)
	}
}

func heroStand(pos *def.Pos){
	curHero.Pos = *pos
	fmt.Println( *curHero )
}

func heroGet(pos *def.Pos){
	
}

func heroDrop(pos *def.Pos){

}
