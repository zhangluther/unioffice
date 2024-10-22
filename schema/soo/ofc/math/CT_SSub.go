// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package math

import (
	"encoding/xml"

	"github.com/zhangluther/unioffice"
)

type CT_SSub struct {
	SSubPr *CT_SSubPr
	E      *CT_OMathArg
	Sub    *CT_OMathArg
}

func NewCT_SSub() *CT_SSub {
	ret := &CT_SSub{}
	ret.E = NewCT_OMathArg()
	ret.Sub = NewCT_OMathArg()
	return ret
}

func (m *CT_SSub) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.SSubPr != nil {
		sesSubPr := xml.StartElement{Name: xml.Name{Local: "m:sSubPr"}}
		e.EncodeElement(m.SSubPr, sesSubPr)
	}
	see := xml.StartElement{Name: xml.Name{Local: "m:e"}}
	e.EncodeElement(m.E, see)
	sesub := xml.StartElement{Name: xml.Name{Local: "m:sub"}}
	e.EncodeElement(m.Sub, sesub)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SSub) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.E = NewCT_OMathArg()
	m.Sub = NewCT_OMathArg()
lCT_SSub:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sSubPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "sSubPr"}:
				m.SSubPr = NewCT_SSubPr()
				if err := d.DecodeElement(m.SSubPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "e"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "e"}:
				if err := d.DecodeElement(m.E, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sub"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "sub"}:
				if err := d.DecodeElement(m.Sub, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_SSub %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SSub
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SSub and its children
func (m *CT_SSub) Validate() error {
	return m.ValidateWithPath("CT_SSub")
}

// ValidateWithPath validates the CT_SSub and its children, prefixing error messages with path
func (m *CT_SSub) ValidateWithPath(path string) error {
	if m.SSubPr != nil {
		if err := m.SSubPr.ValidateWithPath(path + "/SSubPr"); err != nil {
			return err
		}
	}
	if err := m.E.ValidateWithPath(path + "/E"); err != nil {
		return err
	}
	if err := m.Sub.ValidateWithPath(path + "/Sub"); err != nil {
		return err
	}
	return nil
}
