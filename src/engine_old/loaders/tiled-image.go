package loaders

import (
	"encoding/xml"
)

type tsxImage struct {
	XMLName xml.Name `xml:"image"`
	Source  string   `xml:"source,attr"`
}
