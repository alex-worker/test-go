package parser

import "encoding/xml"

type TsxTileSet struct {
	XMLName    xml.Name   `xml:"tileset"`
	Name       string     `xml:"name,attr"`
	TileWidth  string     `xml:"tilewidth,attr"`
	TileHeight string     `xml:"tileheight,attr"`
	Source     string     `xml:"source,attr"`
	Image      tsxImage   `xml:"image"`
	Tiles      []*tsxTile `xml:"tile"`
}
