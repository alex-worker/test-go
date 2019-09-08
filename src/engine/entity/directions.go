package entity

import (
	"../def"
	"errors"
)

func calcNewPos(pos *def.Pos, dir def.Direction, maxPosX uint32, maxPosY uint32) (*def.Pos, error) {

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

	newPos := def.Pos{X: x, Y: y}
	return &newPos, nil

}
