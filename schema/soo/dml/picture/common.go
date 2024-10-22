// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package picture

import "github.com/zhangluther/unioffice"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/picture", "CT_PictureNonVisual", NewCT_PictureNonVisual)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/picture", "CT_Picture", NewCT_Picture)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/picture", "pic", NewPic)
}
