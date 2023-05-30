package entity

import (
	"fmt"
	"test-go/src/engine/defines"
)

// var layers *map[string]*def.Layer
// var myMap *def.Map
var mapW uint32 // размер карты
var mapH uint32

var curHero *defines.Hero

// SetMap устанавливаем карту
func SetMap(mymap *defines.Map) {

	//myMap = mymap
	mapW = mymap.W
	mapH = mymap.H

}

// SetHero устанавливаем главного героя
func SetHero(hero *defines.Hero) {
	curHero = hero
}

// GetHero указатель на героя
func GetHero() *defines.Hero {
	return curHero
}

// HeroDo герой что-то делает
func HeroDo(dir defines.Direction, act defines.HeroAction) {

	oldPos := curHero.Pos

	newPos, err := calcNewPos(&oldPos, dir, mapW, mapH)
	if err != nil {
		fmt.Print(err)
		return
	}

	heroAction(newPos, defines.ActionStand)

}

func heroAction(pos *defines.Pos, act defines.HeroAction) {
	if act == defines.ActionStand {
		heroStand(pos)
	}
}

func heroStand(pos *defines.Pos) {
	curHero.Pos = *pos
	fmt.Println(*curHero)
}

func heroGet(pos *defines.Pos) {

}

func heroDrop(pos *defines.Pos) {

}
