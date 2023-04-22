package TileMap

// TileSetInfo описание тайлсета для дальнейшей обработки
type TileSetInfo struct {
	ImageFileName string
	Tiles         *[]AnimateTile
	TileW         uint64
	TileH         uint64
}
