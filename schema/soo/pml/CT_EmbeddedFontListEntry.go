// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml

import (
	"encoding/xml"

	"github.com/zhangluther/unioffice"
	"github.com/zhangluther/unioffice/schema/soo/dml"
)

type CT_EmbeddedFontListEntry struct {
	// Embedded Font Name
	Font *dml.CT_TextFont
	// Regular Embedded Font
	Regular *CT_EmbeddedFontDataId
	// Bold Embedded Font
	Bold *CT_EmbeddedFontDataId
	// Italic Embedded Font
	Italic *CT_EmbeddedFontDataId
	// Bold Italic Embedded Font
	BoldItalic *CT_EmbeddedFontDataId
}

func NewCT_EmbeddedFontListEntry() *CT_EmbeddedFontListEntry {
	ret := &CT_EmbeddedFontListEntry{}
	ret.Font = dml.NewCT_TextFont()
	return ret
}

func (m *CT_EmbeddedFontListEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sefont := xml.StartElement{Name: xml.Name{Local: "p:font"}}
	e.EncodeElement(m.Font, sefont)
	if m.Regular != nil {
		seregular := xml.StartElement{Name: xml.Name{Local: "p:regular"}}
		e.EncodeElement(m.Regular, seregular)
	}
	if m.Bold != nil {
		sebold := xml.StartElement{Name: xml.Name{Local: "p:bold"}}
		e.EncodeElement(m.Bold, sebold)
	}
	if m.Italic != nil {
		seitalic := xml.StartElement{Name: xml.Name{Local: "p:italic"}}
		e.EncodeElement(m.Italic, seitalic)
	}
	if m.BoldItalic != nil {
		seboldItalic := xml.StartElement{Name: xml.Name{Local: "p:boldItalic"}}
		e.EncodeElement(m.BoldItalic, seboldItalic)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_EmbeddedFontListEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Font = dml.NewCT_TextFont()
lCT_EmbeddedFontListEntry:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "font"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "font"}:
				if err := d.DecodeElement(m.Font, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "regular"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "regular"}:
				m.Regular = NewCT_EmbeddedFontDataId()
				if err := d.DecodeElement(m.Regular, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "bold"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "bold"}:
				m.Bold = NewCT_EmbeddedFontDataId()
				if err := d.DecodeElement(m.Bold, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "italic"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "italic"}:
				m.Italic = NewCT_EmbeddedFontDataId()
				if err := d.DecodeElement(m.Italic, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "boldItalic"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "boldItalic"}:
				m.BoldItalic = NewCT_EmbeddedFontDataId()
				if err := d.DecodeElement(m.BoldItalic, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_EmbeddedFontListEntry %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_EmbeddedFontListEntry
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_EmbeddedFontListEntry and its children
func (m *CT_EmbeddedFontListEntry) Validate() error {
	return m.ValidateWithPath("CT_EmbeddedFontListEntry")
}

// ValidateWithPath validates the CT_EmbeddedFontListEntry and its children, prefixing error messages with path
func (m *CT_EmbeddedFontListEntry) ValidateWithPath(path string) error {
	if err := m.Font.ValidateWithPath(path + "/Font"); err != nil {
		return err
	}
	if m.Regular != nil {
		if err := m.Regular.ValidateWithPath(path + "/Regular"); err != nil {
			return err
		}
	}
	if m.Bold != nil {
		if err := m.Bold.ValidateWithPath(path + "/Bold"); err != nil {
			return err
		}
	}
	if m.Italic != nil {
		if err := m.Italic.ValidateWithPath(path + "/Italic"); err != nil {
			return err
		}
	}
	if m.BoldItalic != nil {
		if err := m.BoldItalic.ValidateWithPath(path + "/BoldItalic"); err != nil {
			return err
		}
	}
	return nil
}
