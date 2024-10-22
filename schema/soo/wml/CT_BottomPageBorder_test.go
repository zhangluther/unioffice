// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/zhangluther/unioffice/schema/soo/wml"
)

func TestCT_BottomPageBorderConstructor(t *testing.T) {
	v := wml.NewCT_BottomPageBorder()
	if v == nil {
		t.Errorf("wml.NewCT_BottomPageBorder must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_BottomPageBorder should validate: %s", err)
	}
}

func TestCT_BottomPageBorderMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_BottomPageBorder()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_BottomPageBorder()
	xml.Unmarshal(buf, v2)
}
