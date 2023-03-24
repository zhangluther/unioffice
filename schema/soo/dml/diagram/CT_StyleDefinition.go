// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package diagram

import (
	"encoding/xml"
	"fmt"

	"github.com/Esword618/unioffice"
	"github.com/Esword618/unioffice/schema/soo/dml"
)

type CT_StyleDefinition struct {
	UniqueIdAttr *string
	MinVerAttr   *string
	Title        []*CT_SDName
	Desc         []*CT_SDDescription
	CatLst       *CT_SDCategories
	Scene3d      *dml.CT_Scene3D
	StyleLbl     []*CT_StyleLabel
	ExtLst       *dml.CT_OfficeArtExtensionList
}

func NewCT_StyleDefinition() *CT_StyleDefinition {
	ret := &CT_StyleDefinition{}
	return ret
}

func (m *CT_StyleDefinition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.UniqueIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uniqueId"},
			Value: fmt.Sprintf("%v", *m.UniqueIdAttr)})
	}
	if m.MinVerAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minVer"},
			Value: fmt.Sprintf("%v", *m.MinVerAttr)})
	}
	e.EncodeToken(start)
	if m.Title != nil {
		setitle := xml.StartElement{Name: xml.Name{Local: "title"}}
		for _, c := range m.Title {
			e.EncodeElement(c, setitle)
		}
	}
	if m.Desc != nil {
		sedesc := xml.StartElement{Name: xml.Name{Local: "desc"}}
		for _, c := range m.Desc {
			e.EncodeElement(c, sedesc)
		}
	}
	if m.CatLst != nil {
		secatLst := xml.StartElement{Name: xml.Name{Local: "catLst"}}
		e.EncodeElement(m.CatLst, secatLst)
	}
	if m.Scene3d != nil {
		sescene3d := xml.StartElement{Name: xml.Name{Local: "scene3d"}}
		e.EncodeElement(m.Scene3d, sescene3d)
	}
	sestyleLbl := xml.StartElement{Name: xml.Name{Local: "styleLbl"}}
	for _, c := range m.StyleLbl {
		e.EncodeElement(c, sestyleLbl)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_StyleDefinition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "uniqueId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UniqueIdAttr = &parsed
			continue
		}
		if attr.Name.Local == "minVer" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MinVerAttr = &parsed
			continue
		}
	}
lCT_StyleDefinition:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "title"}:
				tmp := NewCT_SDName()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Title = append(m.Title, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "desc"}:
				tmp := NewCT_SDDescription()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Desc = append(m.Desc, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "catLst"}:
				m.CatLst = NewCT_SDCategories()
				if err := d.DecodeElement(m.CatLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "scene3d"}:
				m.Scene3d = dml.NewCT_Scene3D()
				if err := d.DecodeElement(m.Scene3d, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "styleLbl"}:
				tmp := NewCT_StyleLabel()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.StyleLbl = append(m.StyleLbl, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "extLst"}:
				m.ExtLst = dml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_StyleDefinition %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_StyleDefinition
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_StyleDefinition and its children
func (m *CT_StyleDefinition) Validate() error {
	return m.ValidateWithPath("CT_StyleDefinition")
}

// ValidateWithPath validates the CT_StyleDefinition and its children, prefixing error messages with path
func (m *CT_StyleDefinition) ValidateWithPath(path string) error {
	for i, v := range m.Title {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Title[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Desc {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Desc[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.CatLst != nil {
		if err := m.CatLst.ValidateWithPath(path + "/CatLst"); err != nil {
			return err
		}
	}
	if m.Scene3d != nil {
		if err := m.Scene3d.ValidateWithPath(path + "/Scene3d"); err != nil {
			return err
		}
	}
	for i, v := range m.StyleLbl {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/StyleLbl[%d]", path, i)); err != nil {
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
