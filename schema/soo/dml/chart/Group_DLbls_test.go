// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/zhangluther/unioffice/schema/soo/dml/chart"
)

func TestGroup_DLblsConstructor(t *testing.T) {
	v := chart.NewGroup_DLbls()
	if v == nil {
		t.Errorf("chart.NewGroup_DLbls must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.Group_DLbls should validate: %s", err)
	}
}

func TestGroup_DLblsMarshalUnmarshal(t *testing.T) {
	v := chart.NewGroup_DLbls()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewGroup_DLbls()
	xml.Unmarshal(buf, v2)
}
