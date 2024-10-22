// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/zhangluther/unioffice/schema/soo/dml"
)

func TestCT_GvmlShapeConstructor(t *testing.T) {
	v := dml.NewCT_GvmlShape()
	if v == nil {
		t.Errorf("dml.NewCT_GvmlShape must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GvmlShape should validate: %s", err)
	}
}

func TestCT_GvmlShapeMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GvmlShape()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GvmlShape()
	xml.Unmarshal(buf, v2)
}
