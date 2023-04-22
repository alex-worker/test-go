package TileMap

// TileSetInfo описание тайлсета для дальнейшей обработки
type TileSetInfo struct {
	Filename string
	Tiles    *map[Cell]*AnimateTile
	TileW    int32
	TileH    int32
}
