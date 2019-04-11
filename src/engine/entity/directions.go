package entity

import (
	"../def"
	"errors"
)

func calcNewPos( pos *def.Pos, dir def.Direction )( *def.Pos, error){

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
		return nil, errors.New("Туда нельзя")
	}

	if ( y<0 ) {
		return nil, errors.New("Туда нельзя")
	}

	newPos := def.Pos{ X: x, Y: y}
	return &newPos, nil

}
