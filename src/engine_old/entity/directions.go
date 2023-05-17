package entity

import (
	"errors"
	"test-go/src/engine/defines"
)

func calcNewPos(pos *defines.Pos, dir defines.Direction, maxPosX uint32, maxPosY uint32) (*defines.Pos, error) {

	x := pos.X
	y := pos.Y

	if dir == defines.DirLeft {
		x--
	} else if dir == defines.DirRight {
		x++
	} else if dir == defines.DirUp {
		y--
	} else if dir == defines.DirDown {
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

	newPos := defines.Pos{X: x, Y: y}
	return &newPos, nil

}
