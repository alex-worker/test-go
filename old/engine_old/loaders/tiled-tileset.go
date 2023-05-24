package loaders

import (
	"encoding/xml"
)

type tsxFrame struct {
	XMLName  xml.Name `xml:"frame"`
	Tileid   uint32   `xml:"tileid,attr"`
	Duration uint32   `xml:"duration,attr"`
}

type tsxAnimation struct {
	XMLName xml.Name    `xml:"animation"`
	Frames  []*tsxFrame `xml:"frame"`
}

type tsxTile struct {
	XMLName    xml.Name     `xml:"tile"`
	ID         uint32       `xml:"id,attr"`
	Animations tsxAnimation `xml:"animation"`
}

type tsxTileSet struct {
	XMLName    xml.Name   `xml:"tileset"`
	Name       string     `xml:"name,attr"`
	TileWidth  string     `xml:"tilewidth,attr"`
	TileHeight string     `xml:"tileheight,attr"`
	Source     string     `xml:"source,attr"`
	Image      tsxImage   `xml:"image"`
	Tiles      []*tsxTile `xml:"tile"`
}
