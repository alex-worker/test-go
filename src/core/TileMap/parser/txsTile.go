package parser

import "encoding/xml"

type tsxTile struct {
	XMLName    xml.Name     `xml:"tile"`
	ID         uint32       `xml:"id,attr"`
	Animations tsxAnimation `xml:"animation"`
}
