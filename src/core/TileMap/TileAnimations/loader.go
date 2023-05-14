package TileAnimations

import (
	. "test-go/src/core/TileMap"
	. "test-go/src/core/TileMap/parser"
)

func convertFrames(anims *TsxAnimation) []AnimateFrame {
	lenFrames := len(anims.Frames)

	frames := make([]AnimateFrame, lenFrames)

	for i, f := range anims.Frames {
		frames[i] = AnimateFrame{
			Cell:     Cell(f.Tileid),
			Duration: f.Duration,
		}
	}
	return frames
}

func convertTileSet(set *TsxTileSet) (string, *TileSet) {
	var tsxFileName = set.Source
	if tsxFileName != "" {
		panic("can't implement file download")
	}

	w, err := StrToUint(set.TileWidth)
	if err != nil {
		panic(err)
	}

	h, err := StrToUint(set.TileHeight)
	if err != nil {
		panic(err)
	}

	fileName := set.Image.Source

	tiles := make(map[Cell]AnimateTile)

	for _, tile := range set.Tiles {
		idCell := Cell(tile.ID)
		//fmt.Println("ID:", idCell)

		frames := convertFrames(&tile.Animations)

		tiles[idCell] = AnimateTile{
			Tick:       0,
			Index:      0,
			NeedUpdate: false,
			Frames:     frames,
		}
	}

	return fileName, &TileSet{
		Tiles: tiles,
		TileW: w,
		TileH: h,
	}
}

type TileSetPack struct {
	FileName string
	Tiles    *TileSet
}

func LoadTileSets(m *TmxMap) []TileSetPack {
	lenTileSets := len(m.TileSets)
	tileSets := make([]TileSetPack, lenTileSets)

	for i, tsxTileSet := range m.TileSets {
		fileName, curTileSet := convertTileSet(tsxTileSet)
		tileSets[i] = TileSetPack{
			FileName: fileName,
			Tiles:    curTileSet,
		}
	}
	return tileSets
}
