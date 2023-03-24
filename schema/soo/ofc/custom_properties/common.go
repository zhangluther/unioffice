// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package custom_properties

import "github.com/Esword618/unioffice"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/custom-properties", "CT_Properties", NewCT_Properties)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/custom-properties", "CT_Property", NewCT_Property)
	unioffice.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/custom-properties", "Properties", NewProperties)
}
