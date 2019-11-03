// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package poly_test

import (
	"testing"

	"github.com/attilaolah/math/poly"
)

func TestInt64TString(t *testing.T) {
	for _, row := range []struct {
		t poly.Int64T
		s string
	}{
		{poly.Int64T{}, "0"},
		{poly.Int64T{poly.Ind{}, 0}, "0"},
		{poly.Int64T{poly.Ind{1, 2, 3}, 0}, "0"},
		{poly.Int64T{poly.Ind{1, 2, 3}, 10}, "10xy²z³"},
		{poly.Int64T{poly.Ind{-1, 0, 1}, -20}, "-20x¯¹z"},
		{poly.Int64T{poly.Ind{}, 5}, "5"},
		{poly.Int64T{poly.Ind{0, 0}, -8}, "-8"},
		{poly.Int64T{poly.Ind{0, 0, 1}, 1}, "z"},
		{poly.Int64T{poly.Ind{0, 1, 1}, 1}, "yz"},
	} {
		if got, want := row.t.String(), row.s; got != want {
			t.Errorf("(%#v).String() = %q; want %q", row.t, got, want)
		}
	}
}

func TestIndString(t *testing.T) {
	for _, row := range []struct {
		i poly.Ind
		s string
	}{
		{poly.Ind{}, "1"},
		{poly.Ind{0}, "1"},
		{poly.Ind{1}, "x"},
		{poly.Ind{1, 1}, "xy"},
		{poly.Ind{1, 0, 1}, "xz"},
		{poly.Ind{0, 1, 0}, "y"},
		{poly.Ind{0, 0, 0, 0}, "1"},
		{poly.Ind{0, 0, 0, 1}, "x₃"},
		{poly.Ind{5}, "x⁵"},
		{poly.Ind{-6}, "x¯⁶"},
		{poly.Ind{123, -321, 0}, "x¹²³y¯³²¹"},
		{poly.Ind{-1, 2, 3, -4, 5, 6}, "x₀¯¹x₁²x₂³x₃¯⁴x₄⁵x₅⁶"},
	} {
		if got, want := row.i.String(), row.s; got != want {
			t.Errorf("(%#v).String() = %q; want %q", row.i, got, want)
		}
	}
}
