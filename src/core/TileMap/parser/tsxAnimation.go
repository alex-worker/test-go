package parser

import "encoding/xml"

type tsxAnimation struct {
	XMLName xml.Name    `xml:"animation"`
	Frames  []*tsxFrame `xml:"frame"`
}
