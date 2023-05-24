package entity

import (
	"errors"
	"test-go/src/engine/def"
	. "test-go/src/math"
)

func calcNewPos(pos *Pos2D, dir def.Direction, maxPosX uint32, maxPosY uint32) (*Pos2D, error) {

	x := pos.X
	y := pos.Y

	if dir == def.DirLeft {
		x--
	} else if dir == def.DirRight {
		x++
	} else if dir == def.DirUp {
		y--
	} else if dir == def.DirDown {
		y++
	}

	if x < 0 {
		return nil, errors.New("туда нельзя")
	}

	if y < 0 {
		return nil, errors.New("туда нельзя")
	}

	if x >= maxPosX {
		return nil, errors.New("туда нельзя")
	}

	if y >= maxPosY {
		return nil, errors.New("туда нельзя")
	}

	newPos := Pos2D{X: x, Y: y}
	return &newPos, nil

}
