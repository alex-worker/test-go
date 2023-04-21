package TileMap

// Layer слой
type Layer struct {
	Data *[][]Cell
	Name string
	W    uint32
	H    uint32
}
