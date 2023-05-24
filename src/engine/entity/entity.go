package entity

import (
	"fmt"
	"test-go/src/engine/def"
	. "test-go/src/math"
)

// var layers *map[string]*def.Layer
// var myMap *def.Map
var mapW uint32 // размер карты
var mapH uint32

var curHero *def.Hero

// SetMap устанавливаем карту
func SetMap(mymap *def.Map) {

	//myMap = mymap
	mapW = mymap.W
	mapH = mymap.H

}

// SetHero устанавливаем главного героя
func SetHero(hero *def.Hero) {
	curHero = hero
}

// GetHero указатель на героя
func GetHero() *def.Hero {
	return curHero
}

// HeroDo герой что-то делает
func HeroDo(dir def.Direction, act def.HeroAction) {

	oldPos := curHero.Pos

	newPos, err := calcNewPos(&oldPos, dir, mapW, mapH)
	if err != nil {
		fmt.Print(err)
		return
	}

	heroAction(newPos, def.ActionStand)

}

func heroAction(pos *Pos2D, act def.HeroAction) {
	if act == def.ActionStand {
		heroStand(pos)
	}
}

func heroStand(pos *Pos2D) {
	curHero.Pos = *pos
	fmt.Println(*curHero)
}

func heroGet(pos *Pos2D) {

}

func heroDrop(pos *Pos2D) {

}
