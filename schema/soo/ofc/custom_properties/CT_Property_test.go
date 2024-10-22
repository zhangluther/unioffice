// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package custom_properties_test

import (
	"encoding/xml"
	"testing"

	"github.com/zhangluther/unioffice/schema/soo/ofc/custom_properties"
)

func TestCT_PropertyConstructor(t *testing.T) {
	v := custom_properties.NewCT_Property()
	if v == nil {
		t.Errorf("custom_properties.NewCT_Property must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed custom_properties.CT_Property should validate: %s", err)
	}
}

func TestCT_PropertyMarshalUnmarshal(t *testing.T) {
	v := custom_properties.NewCT_Property()
	buf, _ := xml.Marshal(v)
	v2 := custom_properties.NewCT_Property()
	xml.Unmarshal(buf, v2)
}
