// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package diagram_test

import (
	"encoding/xml"
	"testing"

	"github.com/zhangluther/unioffice/schema/soo/dml/diagram"
)

func TestCT_CategoryConstructor(t *testing.T) {
	v := diagram.NewCT_Category()
	if v == nil {
		t.Errorf("diagram.NewCT_Category must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Category should validate: %s", err)
	}
}

func TestCT_CategoryMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Category()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Category()
	xml.Unmarshal(buf, v2)
}
