package loaders

import (
	"encoding/xml"
)

type tsxTileSet struct {
	XMLName xml.Name `xml:"tileset"`
	Name string `xml:"name,attr"`
	Width string `xml:"tilewidth,attr"`
	Height string `xml:"tileheight,attr"`	
	Source string `xml:"source,attr"`
	Image tsxImage `xml:"image"`
}
