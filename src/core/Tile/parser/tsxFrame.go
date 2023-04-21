package parser

import "encoding/xml"

type tsxFrame struct {
	XMLName  xml.Name `xml:"frame"`
	Tileid   uint32   `xml:"tileid,attr"`
	Duration uint32   `xml:"duration,attr"`
}
