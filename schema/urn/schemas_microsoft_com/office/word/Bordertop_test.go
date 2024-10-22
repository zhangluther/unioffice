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

func TestBordertopConstructor(t *testing.T) {
	v := word.NewBordertop()
	if v == nil {
		t.Errorf("word.NewBordertop must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed word.Bordertop should validate: %s", err)
	}
}

func TestBordertopMarshalUnmarshal(t *testing.T) {
	v := word.NewBordertop()
	buf, _ := xml.Marshal(v)
	v2 := word.NewBordertop()
	xml.Unmarshal(buf, v2)
}
