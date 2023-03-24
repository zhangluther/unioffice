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

	"github.com/Esword618/unioffice"
)

type CT_BlendEffect struct {
	BlendAttr ST_BlendMode
	Cont      *CT_EffectContainer
}

func NewCT_BlendEffect() *CT_BlendEffect {
	ret := &CT_BlendEffect{}
	ret.BlendAttr = ST_BlendMode(1)
	ret.Cont = NewCT_EffectContainer()
	return ret
}

func (m *CT_BlendEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.BlendAttr.MarshalXMLAttr(xml.Name{Local: "blend"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	secont := xml.StartElement{Name: xml.Name{Local: "a:cont"}}
	e.EncodeElement(m.Cont, secont)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BlendEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.BlendAttr = ST_BlendMode(1)
	m.Cont = NewCT_EffectContainer()
	for _, attr := range start.Attr {
		if attr.Name.Local == "blend" {
			m.BlendAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_BlendEffect:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "cont"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "cont"}:
				if err := d.DecodeElement(m.Cont, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_BlendEffect %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BlendEffect
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BlendEffect and its children
func (m *CT_BlendEffect) Validate() error {
	return m.ValidateWithPath("CT_BlendEffect")
}

// ValidateWithPath validates the CT_BlendEffect and its children, prefixing error messages with path
func (m *CT_BlendEffect) ValidateWithPath(path string) error {
	if m.BlendAttr == ST_BlendModeUnset {
		return fmt.Errorf("%s/BlendAttr is a mandatory field", path)
	}
	if err := m.BlendAttr.ValidateWithPath(path + "/BlendAttr"); err != nil {
		return err
	}
	if err := m.Cont.ValidateWithPath(path + "/Cont"); err != nil {
		return err
	}
	return nil
}
