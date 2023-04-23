package TileMap

// TileSet описание тайлсета для дальнейшей обработки
type TileSet struct {
	ImageFileName string
	Tiles         []AnimateTile
	TileW         uint64
	TileH         uint64
}
