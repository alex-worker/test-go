package TileMap

import "strconv"

func StrToUint(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}
