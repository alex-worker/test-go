package parser

import "encoding/xml"

type TsxAnimation struct {
	XMLName xml.Name    `xml:"animation"`
	Frames  []*tsxFrame `xml:"frame"`
}
