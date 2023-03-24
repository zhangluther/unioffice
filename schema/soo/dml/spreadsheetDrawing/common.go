// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package spreadsheetDrawing

import (
	"encoding/xml"
	"fmt"

	"github.com/Esword618/unioffice"
)

func b2i(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

type ST_EditAs byte

const (
	ST_EditAsUnset    ST_EditAs = 0
	ST_EditAsTwoCell  ST_EditAs = 1
	ST_EditAsOneCell  ST_EditAs = 2
	ST_EditAsAbsolute ST_EditAs = 3
)

func (e ST_EditAs) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_EditAsUnset:
		attr.Value = ""
	case ST_EditAsTwoCell:
		attr.Value = "twoCell"
	case ST_EditAsOneCell:
		attr.Value = "oneCell"
	case ST_EditAsAbsolute:
		attr.Value = "absolute"
	}
	return attr, nil
}

func (e *ST_EditAs) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "twoCell":
		*e = 1
	case "oneCell":
		*e = 2
	case "absolute":
		*e = 3
	}
	return nil
}

func (m ST_EditAs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_EditAs) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "twoCell":
			*m = 1
		case "oneCell":
			*m = 2
		case "absolute":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_EditAs) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "twoCell"
	case 2:
		return "oneCell"
	case 3:
		return "absolute"
	}
	return ""
}

func (m ST_EditAs) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_EditAs) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "CT_AnchorClientData", NewCT_AnchorClientData)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "CT_ShapeNonVisual", NewCT_ShapeNonVisual)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "CT_Shape", NewCT_Shape)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "CT_ConnectorNonVisual", NewCT_ConnectorNonVisual)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "CT_Connector", NewCT_Connector)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "CT_PictureNonVisual", NewCT_PictureNonVisual)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "CT_Picture", NewCT_Picture)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "CT_GraphicalObjectFrameNonVisual", NewCT_GraphicalObjectFrameNonVisual)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "CT_GraphicalObjectFrame", NewCT_GraphicalObjectFrame)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "CT_GroupShapeNonVisual", NewCT_GroupShapeNonVisual)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "CT_GroupShape", NewCT_GroupShape)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "CT_Rel", NewCT_Rel)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "CT_Marker", NewCT_Marker)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "CT_TwoCellAnchor", NewCT_TwoCellAnchor)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "CT_OneCellAnchor", NewCT_OneCellAnchor)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "CT_AbsoluteAnchor", NewCT_AbsoluteAnchor)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "CT_Drawing", NewCT_Drawing)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "from", NewFrom)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "to", NewTo)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "wsDr", NewWsDr)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "EG_ObjectChoices", NewEG_ObjectChoices)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", "EG_Anchor", NewEG_Anchor)
}
