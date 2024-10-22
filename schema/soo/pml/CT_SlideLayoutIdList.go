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
	"fmt"

	"github.com/zhangluther/unioffice"
)

type CT_SlideLayoutIdList struct {
	// Slide Layout Id
	SldLayoutId []*CT_SlideLayoutIdListEntry
}

func NewCT_SlideLayoutIdList() *CT_SlideLayoutIdList {
	ret := &CT_SlideLayoutIdList{}
	return ret
}

func (m *CT_SlideLayoutIdList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.SldLayoutId != nil {
		sesldLayoutId := xml.StartElement{Name: xml.Name{Local: "p:sldLayoutId"}}
		for _, c := range m.SldLayoutId {
			e.EncodeElement(c, sesldLayoutId)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SlideLayoutIdList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SlideLayoutIdList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "sldLayoutId"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "sldLayoutId"}:
				tmp := NewCT_SlideLayoutIdListEntry()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SldLayoutId = append(m.SldLayoutId, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_SlideLayoutIdList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SlideLayoutIdList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SlideLayoutIdList and its children
func (m *CT_SlideLayoutIdList) Validate() error {
	return m.ValidateWithPath("CT_SlideLayoutIdList")
}

// ValidateWithPath validates the CT_SlideLayoutIdList and its children, prefixing error messages with path
func (m *CT_SlideLayoutIdList) ValidateWithPath(path string) error {
	for i, v := range m.SldLayoutId {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SldLayoutId[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
