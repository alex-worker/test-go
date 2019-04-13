package entity

import (
	"fmt"
	"../def"
)

var layers *map[string]*def.Layer
var mapW uint32 // размер карты
var mapH uint32

var curHero *def.Hero

// SetMap устанавливаем карту
func SetMap(layersList *map[string]*def.Layer){

	layers = layersList
	
	var firstLayer *def.Layer
	for _, firstLayer = range *layers{
		break
	}

	data := *firstLayer.Data
	mapW = uint32(len(data))
	mapH = uint32(len(data[0]))

}

// SetHero устанавливаем главного героя
func SetHero( hero *def.Hero ){
	curHero = hero
}

// GetMap получаем карту
func GetMap() ( layerList *map[string]*def.Layer, mapW int, mapH int){
	return layers, mapW, mapH
}

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
