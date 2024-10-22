// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package vml

import (
	"encoding/xml"
	"fmt"

	"github.com/zhangluther/unioffice/schema/soo/ofc/sharedTypes"
)

type CT_Shadow struct {
	OnAttr       sharedTypes.ST_TrueFalse
	TypeAttr     ST_ShadowType
	ObscuredAttr sharedTypes.ST_TrueFalse
	ColorAttr    *string
	OpacityAttr  *string
	OffsetAttr   *string
	Color2Attr   *string
	Offset2Attr  *string
	OriginAttr   *string
	MatrixAttr   *string
	IdAttr       *string
}

func NewCT_Shadow() *CT_Shadow {
	ret := &CT_Shadow{}
	return ret
}

func (m *CT_Shadow) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.OnAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.OnAttr.MarshalXMLAttr(xml.Name{Local: "on"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.TypeAttr != ST_ShadowTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ObscuredAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.ObscuredAttr.MarshalXMLAttr(xml.Name{Local: "obscured"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ColorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "color"},
			Value: fmt.Sprintf("%v", *m.ColorAttr)})
	}
	if m.OpacityAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "opacity"},
			Value: fmt.Sprintf("%v", *m.OpacityAttr)})
	}
	if m.OffsetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "offset"},
			Value: fmt.Sprintf("%v", *m.OffsetAttr)})
	}
	if m.Color2Attr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "color2"},
			Value: fmt.Sprintf("%v", *m.Color2Attr)})
	}
	if m.Offset2Attr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "offset2"},
			Value: fmt.Sprintf("%v", *m.Offset2Attr)})
	}
	if m.OriginAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "origin"},
			Value: fmt.Sprintf("%v", *m.OriginAttr)})
	}
	if m.MatrixAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "matrix"},
			Value: fmt.Sprintf("%v", *m.MatrixAttr)})
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Shadow) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "on" {
			m.OnAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "obscured" {
			m.ObscuredAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "color" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ColorAttr = &parsed
			continue
		}
		if attr.Name.Local == "opacity" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OpacityAttr = &parsed
			continue
		}
		if attr.Name.Local == "offset" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OffsetAttr = &parsed
			continue
		}
		if attr.Name.Local == "color2" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Color2Attr = &parsed
			continue
		}
		if attr.Name.Local == "offset2" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Offset2Attr = &parsed
			continue
		}
		if attr.Name.Local == "origin" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OriginAttr = &parsed
			continue
		}
		if attr.Name.Local == "matrix" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MatrixAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Shadow: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Shadow and its children
func (m *CT_Shadow) Validate() error {
	return m.ValidateWithPath("CT_Shadow")
}

// ValidateWithPath validates the CT_Shadow and its children, prefixing error messages with path
func (m *CT_Shadow) ValidateWithPath(path string) error {
	if err := m.OnAttr.ValidateWithPath(path + "/OnAttr"); err != nil {
		return err
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.ObscuredAttr.ValidateWithPath(path + "/ObscuredAttr"); err != nil {
		return err
	}
	return nil
}
