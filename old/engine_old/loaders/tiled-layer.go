package loaders

import (
	"encoding/xml"
)

type tmxLayer struct {
	XMLName xml.Name `xml:"layer"`
	Name    string   `xml:"name,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	Data    string   `xml:"data"`
}
