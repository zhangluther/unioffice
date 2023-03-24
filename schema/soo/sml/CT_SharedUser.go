// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	"github.com/Esword618/unioffice"
	"github.com/Esword618/unioffice/schema/soo/ofc/sharedTypes"
)

type CT_SharedUser struct {
	// User Revisions GUID
	GuidAttr string
	// User Name
	NameAttr string
	// User Id
	IdAttr int32
	// Date Time
	DateTimeAttr time.Time
	ExtLst       *CT_ExtensionList
}

func NewCT_SharedUser() *CT_SharedUser {
	ret := &CT_SharedUser{}
	ret.GuidAttr = "{00000000-0000-0000-0000-000000000000}"
	return ret
}

func (m *CT_SharedUser) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "guid"},
		Value: fmt.Sprintf("%v", m.GuidAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dateTime"},
		Value: fmt.Sprintf("%v", m.DateTimeAttr)})
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SharedUser) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.GuidAttr = "{00000000-0000-0000-0000-000000000000}"
	for _, attr := range start.Attr {
		if attr.Name.Local == "guid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GuidAttr = parsed
			continue
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdAttr = int32(parsed)
			continue
		}
		if attr.Name.Local == "dateTime" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.DateTimeAttr = parsed
			continue
		}
	}
lCT_SharedUser:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_SharedUser %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SharedUser
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SharedUser and its children
func (m *CT_SharedUser) Validate() error {
	return m.ValidateWithPath("CT_SharedUser")
}

// ValidateWithPath validates the CT_SharedUser and its children, prefixing error messages with path
func (m *CT_SharedUser) ValidateWithPath(path string) error {
	if !sharedTypes.ST_GuidPatternRe.MatchString(m.GuidAttr) {
		return fmt.Errorf(`%s/m.GuidAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, m.GuidAttr)
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
