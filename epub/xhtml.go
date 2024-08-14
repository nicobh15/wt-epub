package epub

import (
	"encoding/xml"
)

type xhtml struct {
	xml *xhtmlDoc
}

type xhtmlDoc struct {
	XMLName   xml.Name `xml:"http://www.w3.org/1999/xhtml html"`
	XmlnsEpub string   `xml:"xmlns:epub,attr,omitempty"`
}
