// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/zhangluther/unioffice"
)

type CT_NonVisualContentPartProperties struct {
	IsCommentAttr *bool
	CpLocks       *CT_ContentPartLocking
	ExtLst        *CT_OfficeArtExtensionList
}

func NewCT_NonVisualContentPartProperties() *CT_NonVisualContentPartProperties {
	ret := &CT_NonVisualContentPartProperties{}
	return ret
}

func (m *CT_NonVisualContentPartProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.IsCommentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "isComment"},
			Value: fmt.Sprintf("%d", b2i(*m.IsCommentAttr))})
	}
	e.EncodeToken(start)
	if m.CpLocks != nil {
		secpLocks := xml.StartElement{Name: xml.Name{Local: "a:cpLocks"}}
		e.EncodeElement(m.CpLocks, secpLocks)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NonVisualContentPartProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "isComment" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.IsCommentAttr = &parsed
			continue
		}
	}
lCT_NonVisualContentPartProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "cpLocks"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "cpLocks"}:
				m.CpLocks = NewCT_ContentPartLocking()
				if err := d.DecodeElement(m.CpLocks, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_NonVisualContentPartProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NonVisualContentPartProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_NonVisualContentPartProperties and its children
func (m *CT_NonVisualContentPartProperties) Validate() error {
	return m.ValidateWithPath("CT_NonVisualContentPartProperties")
}

// ValidateWithPath validates the CT_NonVisualContentPartProperties and its children, prefixing error messages with path
func (m *CT_NonVisualContentPartProperties) ValidateWithPath(path string) error {
	if m.CpLocks != nil {
		if err := m.CpLocks.ValidateWithPath(path + "/CpLocks"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
