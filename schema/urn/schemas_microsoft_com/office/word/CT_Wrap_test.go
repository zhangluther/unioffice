// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package word_test

import (
	"encoding/xml"
	"testing"

	"github.com/zhangluther/unioffice/schema/urn/schemas_microsoft_com/office/word"
)

func TestCT_WrapConstructor(t *testing.T) {
	v := word.NewCT_Wrap()
	if v == nil {
		t.Errorf("word.NewCT_Wrap must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed word.CT_Wrap should validate: %s", err)
	}
}

func TestCT_WrapMarshalUnmarshal(t *testing.T) {
	v := word.NewCT_Wrap()
	buf, _ := xml.Marshal(v)
	v2 := word.NewCT_Wrap()
	xml.Unmarshal(buf, v2)
}
