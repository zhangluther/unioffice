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

	"github.com/zhangluther/unioffice"
	"github.com/zhangluther/unioffice/schema/urn/schemas_microsoft_com/office/excel"
	"github.com/zhangluther/unioffice/schema/urn/schemas_microsoft_com/office/powerpoint"
	"github.com/zhangluther/unioffice/schema/urn/schemas_microsoft_com/office/word"
)

type EG_ShapeElements struct {
	Path          *Path
	Formulas      *Formulas
	Handles       *Handles
	Fill          *Fill
	Stroke        *Stroke
	Shadow        *Shadow
	Textbox       *Textbox
	Textpath      *Textpath
	Imagedata     *Imagedata
	Skew          *OfcSkew
	Extrusion     *OfcExtrusion
	Callout       *OfcCallout
	Lock          *OfcLock
	Clippath      *OfcClippath
	Signatureline *OfcSignatureline
	Wrap          *word.Wrap
	Anchorlock    *word.Anchorlock
	Bordertop     *word.Bordertop
	Borderbottom  *word.Borderbottom
	Borderleft    *word.Borderleft
	Borderright   *word.Borderright
	ClientData    *excel.ClientData
	Textdata      *powerpoint.Textdata
}

func NewEG_ShapeElements() *EG_ShapeElements {
	ret := &EG_ShapeElements{}
	return ret
}

func (m *EG_ShapeElements) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Path != nil {
		sepath := xml.StartElement{Name: xml.Name{Local: "v:path"}}
		e.EncodeElement(m.Path, sepath)
	}
	if m.Formulas != nil {
		seformulas := xml.StartElement{Name: xml.Name{Local: "v:formulas"}}
		e.EncodeElement(m.Formulas, seformulas)
	}
	if m.Handles != nil {
		sehandles := xml.StartElement{Name: xml.Name{Local: "v:handles"}}
		e.EncodeElement(m.Handles, sehandles)
	}
	if m.Fill != nil {
		sefill := xml.StartElement{Name: xml.Name{Local: "v:fill"}}
		e.EncodeElement(m.Fill, sefill)
	}
	if m.Stroke != nil {
		sestroke := xml.StartElement{Name: xml.Name{Local: "v:stroke"}}
		e.EncodeElement(m.Stroke, sestroke)
	}
	if m.Shadow != nil {
		seshadow := xml.StartElement{Name: xml.Name{Local: "v:shadow"}}
		e.EncodeElement(m.Shadow, seshadow)
	}
	if m.Textbox != nil {
		setextbox := xml.StartElement{Name: xml.Name{Local: "v:textbox"}}
		e.EncodeElement(m.Textbox, setextbox)
	}
	if m.Textpath != nil {
		setextpath := xml.StartElement{Name: xml.Name{Local: "v:textpath"}}
		e.EncodeElement(m.Textpath, setextpath)
	}
	if m.Imagedata != nil {
		seimagedata := xml.StartElement{Name: xml.Name{Local: "v:imagedata"}}
		e.EncodeElement(m.Imagedata, seimagedata)
	}
	if m.Skew != nil {
		seskew := xml.StartElement{Name: xml.Name{Local: "o:skew"}}
		e.EncodeElement(m.Skew, seskew)
	}
	if m.Extrusion != nil {
		seextrusion := xml.StartElement{Name: xml.Name{Local: "o:extrusion"}}
		e.EncodeElement(m.Extrusion, seextrusion)
	}
	if m.Callout != nil {
		secallout := xml.StartElement{Name: xml.Name{Local: "o:callout"}}
		e.EncodeElement(m.Callout, secallout)
	}
	if m.Lock != nil {
		selock := xml.StartElement{Name: xml.Name{Local: "o:lock"}}
		e.EncodeElement(m.Lock, selock)
	}
	if m.Clippath != nil {
		seclippath := xml.StartElement{Name: xml.Name{Local: "o:clippath"}}
		e.EncodeElement(m.Clippath, seclippath)
	}
	if m.Signatureline != nil {
		sesignatureline := xml.StartElement{Name: xml.Name{Local: "o:signatureline"}}
		e.EncodeElement(m.Signatureline, sesignatureline)
	}
	if m.Wrap != nil {
		sewrap := xml.StartElement{Name: xml.Name{Local: "urn:wrap"}}
		e.EncodeElement(m.Wrap, sewrap)
	}
	if m.Anchorlock != nil {
		seanchorlock := xml.StartElement{Name: xml.Name{Local: "urn:anchorlock"}}
		e.EncodeElement(m.Anchorlock, seanchorlock)
	}
	if m.Bordertop != nil {
		sebordertop := xml.StartElement{Name: xml.Name{Local: "urn:bordertop"}}
		e.EncodeElement(m.Bordertop, sebordertop)
	}
	if m.Borderbottom != nil {
		seborderbottom := xml.StartElement{Name: xml.Name{Local: "urn:borderbottom"}}
		e.EncodeElement(m.Borderbottom, seborderbottom)
	}
	if m.Borderleft != nil {
		seborderleft := xml.StartElement{Name: xml.Name{Local: "urn:borderleft"}}
		e.EncodeElement(m.Borderleft, seborderleft)
	}
	if m.Borderright != nil {
		seborderright := xml.StartElement{Name: xml.Name{Local: "urn:borderright"}}
		e.EncodeElement(m.Borderright, seborderright)
	}
	if m.ClientData != nil {
		seClientData := xml.StartElement{Name: xml.Name{Local: "x:ClientData"}}
		e.EncodeElement(m.ClientData, seClientData)
	}
	if m.Textdata != nil {
		setextdata := xml.StartElement{Name: xml.Name{Local: "ur:textdata"}}
		e.EncodeElement(m.Textdata, setextdata)
	}
	return nil
}

func (m *EG_ShapeElements) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ShapeElements:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "path"}:
				m.Path = NewPath()
				if err := d.DecodeElement(m.Path, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "formulas"}:
				m.Formulas = NewFormulas()
				if err := d.DecodeElement(m.Formulas, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "handles"}:
				m.Handles = NewHandles()
				if err := d.DecodeElement(m.Handles, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "fill"}:
				m.Fill = NewFill()
				if err := d.DecodeElement(m.Fill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "stroke"}:
				m.Stroke = NewStroke()
				if err := d.DecodeElement(m.Stroke, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "shadow"}:
				m.Shadow = NewShadow()
				if err := d.DecodeElement(m.Shadow, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "textbox"}:
				m.Textbox = NewTextbox()
				if err := d.DecodeElement(m.Textbox, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "textpath"}:
				m.Textpath = NewTextpath()
				if err := d.DecodeElement(m.Textpath, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "imagedata"}:
				m.Imagedata = NewImagedata()
				if err := d.DecodeElement(m.Imagedata, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "skew"}:
				m.Skew = NewOfcSkew()
				if err := d.DecodeElement(m.Skew, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "extrusion"}:
				m.Extrusion = NewOfcExtrusion()
				if err := d.DecodeElement(m.Extrusion, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "callout"}:
				m.Callout = NewOfcCallout()
				if err := d.DecodeElement(m.Callout, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "lock"}:
				m.Lock = NewOfcLock()
				if err := d.DecodeElement(m.Lock, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "clippath"}:
				m.Clippath = NewOfcClippath()
				if err := d.DecodeElement(m.Clippath, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "signatureline"}:
				m.Signatureline = NewOfcSignatureline()
				if err := d.DecodeElement(m.Signatureline, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "wrap"}:
				m.Wrap = word.NewWrap()
				if err := d.DecodeElement(m.Wrap, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "anchorlock"}:
				m.Anchorlock = word.NewAnchorlock()
				if err := d.DecodeElement(m.Anchorlock, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "bordertop"}:
				m.Bordertop = word.NewBordertop()
				if err := d.DecodeElement(m.Bordertop, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "borderbottom"}:
				m.Borderbottom = word.NewBorderbottom()
				if err := d.DecodeElement(m.Borderbottom, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "borderleft"}:
				m.Borderleft = word.NewBorderleft()
				if err := d.DecodeElement(m.Borderleft, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "borderright"}:
				m.Borderright = word.NewBorderright()
				if err := d.DecodeElement(m.Borderright, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ClientData"}:
				m.ClientData = excel.NewClientData()
				if err := d.DecodeElement(m.ClientData, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:powerpoint", Local: "textdata"}:
				m.Textdata = powerpoint.NewTextdata()
				if err := d.DecodeElement(m.Textdata, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on EG_ShapeElements %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_ShapeElements
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ShapeElements and its children
func (m *EG_ShapeElements) Validate() error {
	return m.ValidateWithPath("EG_ShapeElements")
}

// ValidateWithPath validates the EG_ShapeElements and its children, prefixing error messages with path
func (m *EG_ShapeElements) ValidateWithPath(path string) error {
	if m.Path != nil {
		if err := m.Path.ValidateWithPath(path + "/Path"); err != nil {
			return err
		}
	}
	if m.Formulas != nil {
		if err := m.Formulas.ValidateWithPath(path + "/Formulas"); err != nil {
			return err
		}
	}
	if m.Handles != nil {
		if err := m.Handles.ValidateWithPath(path + "/Handles"); err != nil {
			return err
		}
	}
	if m.Fill != nil {
		if err := m.Fill.ValidateWithPath(path + "/Fill"); err != nil {
			return err
		}
	}
	if m.Stroke != nil {
		if err := m.Stroke.ValidateWithPath(path + "/Stroke"); err != nil {
			return err
		}
	}
	if m.Shadow != nil {
		if err := m.Shadow.ValidateWithPath(path + "/Shadow"); err != nil {
			return err
		}
	}
	if m.Textbox != nil {
		if err := m.Textbox.ValidateWithPath(path + "/Textbox"); err != nil {
			return err
		}
	}
	if m.Textpath != nil {
		if err := m.Textpath.ValidateWithPath(path + "/Textpath"); err != nil {
			return err
		}
	}
	if m.Imagedata != nil {
		if err := m.Imagedata.ValidateWithPath(path + "/Imagedata"); err != nil {
			return err
		}
	}
	if m.Skew != nil {
		if err := m.Skew.ValidateWithPath(path + "/Skew"); err != nil {
			return err
		}
	}
	if m.Extrusion != nil {
		if err := m.Extrusion.ValidateWithPath(path + "/Extrusion"); err != nil {
			return err
		}
	}
	if m.Callout != nil {
		if err := m.Callout.ValidateWithPath(path + "/Callout"); err != nil {
			return err
		}
	}
	if m.Lock != nil {
		if err := m.Lock.ValidateWithPath(path + "/Lock"); err != nil {
			return err
		}
	}
	if m.Clippath != nil {
		if err := m.Clippath.ValidateWithPath(path + "/Clippath"); err != nil {
			return err
		}
	}
	if m.Signatureline != nil {
		if err := m.Signatureline.ValidateWithPath(path + "/Signatureline"); err != nil {
			return err
		}
	}
	if m.Wrap != nil {
		if err := m.Wrap.ValidateWithPath(path + "/Wrap"); err != nil {
			return err
		}
	}
	if m.Anchorlock != nil {
		if err := m.Anchorlock.ValidateWithPath(path + "/Anchorlock"); err != nil {
			return err
		}
	}
	if m.Bordertop != nil {
		if err := m.Bordertop.ValidateWithPath(path + "/Bordertop"); err != nil {
			return err
		}
	}
	if m.Borderbottom != nil {
		if err := m.Borderbottom.ValidateWithPath(path + "/Borderbottom"); err != nil {
			return err
		}
	}
	if m.Borderleft != nil {
		if err := m.Borderleft.ValidateWithPath(path + "/Borderleft"); err != nil {
			return err
		}
	}
	if m.Borderright != nil {
		if err := m.Borderright.ValidateWithPath(path + "/Borderright"); err != nil {
			return err
		}
	}
	if m.ClientData != nil {
		if err := m.ClientData.ValidateWithPath(path + "/ClientData"); err != nil {
			return err
		}
	}
	if m.Textdata != nil {
		if err := m.Textdata.ValidateWithPath(path + "/Textdata"); err != nil {
			return err
		}
	}
	return nil
}
