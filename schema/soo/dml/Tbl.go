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

	"github.com/zhangluther/unioffice"
)

type Tbl struct {
	CT_Table
}

func NewTbl() *Tbl {
	ret := &Tbl{}
	ret.CT_Table = *NewCT_Table()
	return ret
}

func (m *Tbl) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "a:tbl"
	return m.CT_Table.MarshalXML(e, start)
}

func (m *Tbl) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Table = *NewCT_Table()
lTbl:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "tblPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "tblPr"}:
				m.TblPr = NewCT_TableProperties()
				if err := d.DecodeElement(m.TblPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "tblGrid"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "tblGrid"}:
				if err := d.DecodeElement(m.TblGrid, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "tr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "tr"}:
				tmp := NewCT_TableRow()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tr = append(m.Tr, tmp)
			default:
				unioffice.Log("skipping unsupported element on Tbl %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lTbl
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Tbl and its children
func (m *Tbl) Validate() error {
	return m.ValidateWithPath("Tbl")
}

// ValidateWithPath validates the Tbl and its children, prefixing error messages with path
func (m *Tbl) ValidateWithPath(path string) error {
	if err := m.CT_Table.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
