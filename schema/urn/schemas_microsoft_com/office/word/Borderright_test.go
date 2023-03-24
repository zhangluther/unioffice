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

	"github.com/Esword618/unioffice/schema/urn/schemas_microsoft_com/office/word"
)

func TestBorderrightConstructor(t *testing.T) {
	v := word.NewBorderright()
	if v == nil {
		t.Errorf("word.NewBorderright must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed word.Borderright should validate: %s", err)
	}
}

func TestBorderrightMarshalUnmarshal(t *testing.T) {
	v := word.NewBorderright()
	buf, _ := xml.Marshal(v)
	v2 := word.NewBorderright()
	xml.Unmarshal(buf, v2)
}
