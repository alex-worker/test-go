package parser

import (
	"encoding/xml"
)

type TmxMap struct {
	XMLName    xml.Name      `xml:"map"`
	Width      string        `xml:"width,attr"`
	Height     string        `xml:"height,attr"`
	TileWidth  string        `xml:"tilewidth,attr"`
	TileHeight string        `xml:"tileheight,attr"`
	TileSets   []*tsxTileSet `xml:"tileset"`
	Layers     []*tmxLayer   `xml:"layer"`
}
