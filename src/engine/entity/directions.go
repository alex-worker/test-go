package entity

import (
	"../def"
)

func calcNewPos( pos *def.Pos, dir def.Direction )( *def.Pos, bool){

	x := pos.X
	y := pos.Y

	if ( dir == def.DirLeft ){
		x--
	} else if ( dir == def.DirRight ){
		x++
	} else if ( dir == def.DirUp ){
		y--
	} else if ( dir == def.DirDown ){
		y++
	}

	if ( x<0 ) {
		return nil, false
	}

	newPos := def.Pos{ X: x, Y: y}
	return &newPos, true

}
