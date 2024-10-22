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

	"github.com/zhangluther/unioffice"
)

type CT_LayoutVariablePropertySet struct {
	OrgChart      *CT_OrgChart
	ChMax         *CT_ChildMax
	ChPref        *CT_ChildPref
	BulletEnabled *CT_BulletEnabled
	Dir           *CT_Direction
	HierBranch    *CT_HierBranchStyle
	AnimOne       *CT_AnimOne
	AnimLvl       *CT_AnimLvl
	ResizeHandles *CT_ResizeHandles
}

func NewCT_LayoutVariablePropertySet() *CT_LayoutVariablePropertySet {
	ret := &CT_LayoutVariablePropertySet{}
	return ret
}

func (m *CT_LayoutVariablePropertySet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.OrgChart != nil {
		seorgChart := xml.StartElement{Name: xml.Name{Local: "orgChart"}}
		e.EncodeElement(m.OrgChart, seorgChart)
	}
	if m.ChMax != nil {
		sechMax := xml.StartElement{Name: xml.Name{Local: "chMax"}}
		e.EncodeElement(m.ChMax, sechMax)
	}
	if m.ChPref != nil {
		sechPref := xml.StartElement{Name: xml.Name{Local: "chPref"}}
		e.EncodeElement(m.ChPref, sechPref)
	}
	if m.BulletEnabled != nil {
		sebulletEnabled := xml.StartElement{Name: xml.Name{Local: "bulletEnabled"}}
		e.EncodeElement(m.BulletEnabled, sebulletEnabled)
	}
	if m.Dir != nil {
		sedir := xml.StartElement{Name: xml.Name{Local: "dir"}}
		e.EncodeElement(m.Dir, sedir)
	}
	if m.HierBranch != nil {
		sehierBranch := xml.StartElement{Name: xml.Name{Local: "hierBranch"}}
		e.EncodeElement(m.HierBranch, sehierBranch)
	}
	if m.AnimOne != nil {
		seanimOne := xml.StartElement{Name: xml.Name{Local: "animOne"}}
		e.EncodeElement(m.AnimOne, seanimOne)
	}
	if m.AnimLvl != nil {
		seanimLvl := xml.StartElement{Name: xml.Name{Local: "animLvl"}}
		e.EncodeElement(m.AnimLvl, seanimLvl)
	}
	if m.ResizeHandles != nil {
		seresizeHandles := xml.StartElement{Name: xml.Name{Local: "resizeHandles"}}
		e.EncodeElement(m.ResizeHandles, seresizeHandles)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LayoutVariablePropertySet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_LayoutVariablePropertySet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "orgChart"}:
				m.OrgChart = NewCT_OrgChart()
				if err := d.DecodeElement(m.OrgChart, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "chMax"}:
				m.ChMax = NewCT_ChildMax()
				if err := d.DecodeElement(m.ChMax, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "chPref"}:
				m.ChPref = NewCT_ChildPref()
				if err := d.DecodeElement(m.ChPref, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "bulletEnabled"}:
				m.BulletEnabled = NewCT_BulletEnabled()
				if err := d.DecodeElement(m.BulletEnabled, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "dir"}:
				m.Dir = NewCT_Direction()
				if err := d.DecodeElement(m.Dir, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "hierBranch"}:
				m.HierBranch = NewCT_HierBranchStyle()
				if err := d.DecodeElement(m.HierBranch, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "animOne"}:
				m.AnimOne = NewCT_AnimOne()
				if err := d.DecodeElement(m.AnimOne, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "animLvl"}:
				m.AnimLvl = NewCT_AnimLvl()
				if err := d.DecodeElement(m.AnimLvl, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "resizeHandles"}:
				m.ResizeHandles = NewCT_ResizeHandles()
				if err := d.DecodeElement(m.ResizeHandles, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_LayoutVariablePropertySet %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_LayoutVariablePropertySet
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_LayoutVariablePropertySet and its children
func (m *CT_LayoutVariablePropertySet) Validate() error {
	return m.ValidateWithPath("CT_LayoutVariablePropertySet")
}

// ValidateWithPath validates the CT_LayoutVariablePropertySet and its children, prefixing error messages with path
func (m *CT_LayoutVariablePropertySet) ValidateWithPath(path string) error {
	if m.OrgChart != nil {
		if err := m.OrgChart.ValidateWithPath(path + "/OrgChart"); err != nil {
			return err
		}
	}
	if m.ChMax != nil {
		if err := m.ChMax.ValidateWithPath(path + "/ChMax"); err != nil {
			return err
		}
	}
	if m.ChPref != nil {
		if err := m.ChPref.ValidateWithPath(path + "/ChPref"); err != nil {
			return err
		}
	}
	if m.BulletEnabled != nil {
		if err := m.BulletEnabled.ValidateWithPath(path + "/BulletEnabled"); err != nil {
			return err
		}
	}
	if m.Dir != nil {
		if err := m.Dir.ValidateWithPath(path + "/Dir"); err != nil {
			return err
		}
	}
	if m.HierBranch != nil {
		if err := m.HierBranch.ValidateWithPath(path + "/HierBranch"); err != nil {
			return err
		}
	}
	if m.AnimOne != nil {
		if err := m.AnimOne.ValidateWithPath(path + "/AnimOne"); err != nil {
			return err
		}
	}
	if m.AnimLvl != nil {
		if err := m.AnimLvl.ValidateWithPath(path + "/AnimLvl"); err != nil {
			return err
		}
	}
	if m.ResizeHandles != nil {
		if err := m.ResizeHandles.ValidateWithPath(path + "/ResizeHandles"); err != nil {
			return err
		}
	}
	return nil
}
