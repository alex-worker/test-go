package entity

import (
	"fmt"
	"../def"
)

// var layers *map[string]*def.Layer
var myMap *def.Map
var mapW uint32 // размер карты
var mapH uint32

var curHero *def.Hero

// SetMap устанавливаем карту
func SetMap(mymap *def.Map){

	myMap = mymap
	

	// mapW = 
	// mapH = uint32(len(data[0]))

	println( mapW, mapH )

}

// SetHero устанавливаем главного героя
func SetHero( hero *def.Hero ){
	curHero = hero
}

// GetMap получаем карту
// func GetMap() ( layerList *map[string]*def.Layer, w uint32, h uint32){
	// return layers, mapW, mapH
// }

// GetHero указатель на героя
func GetHero() *def.Hero {
	return curHero
}
 
// HeroDo герой что-то делает
func HeroDo( dir def.Direction, act def.HeroAction ){

	oldPos := curHero.Pos

	newPos,err := calcNewPos( &oldPos, dir, int(mapW), int(mapH) )
	if err != nil {
		fmt.Print(err)
		return
	}

	heroAction( newPos, def.ActionStand )

}

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
