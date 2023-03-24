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

	"github.com/Esword618/unioffice/schema/soo/dml/chart"
)

func TestCT_Area3DChartConstructor(t *testing.T) {
	v := chart.NewCT_Area3DChart()
	if v == nil {
		t.Errorf("chart.NewCT_Area3DChart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Area3DChart should validate: %s", err)
	}
}

func TestCT_Area3DChartMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Area3DChart()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Area3DChart()
	xml.Unmarshal(buf, v2)
}
