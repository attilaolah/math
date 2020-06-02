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

	"github.com/attilaolah/math/go/poly"
)

func TestInt64PAdd(t *testing.T) {
	for _, row := range []struct {
		a, b poly.Int64P
		s    string
	}{
		{
			poly.Int64P{
				poly.Int64T{poly.Ind{1}, 2},
			},
			poly.Int64P{
				poly.Int64T{poly.Ind{1}, 2},
			},
			"4x",
		},
		{
			poly.Int64P{
				poly.Int64T{poly.Ind{0, 1}, 1},
				poly.Int64T{poly.Ind{1, 0}, 1},
				poly.Int64T{poly.Ind{1, 1}, 2},
			},
			poly.Int64P{
				poly.Int64T{poly.Ind{1, 1}, 2},
			},
			"4xy + x + y",
		},
	} {
		if got, want := row.a.Add(row.b).String(), row.s; got != want {
			t.Errorf("(%q).Add(%q) = %q; want %q", row.a, row.b, got, want)
		}
	}
}

func TestInt64PMul(t *testing.T) {
	for _, row := range []struct {
		a, b poly.Int64P
		s    string
	}{
		{
			poly.Int64P{
				poly.Int64T{poly.Ind{1}, 2},
			},
			poly.Int64P{
				poly.Int64T{poly.Ind{1}, 2},
			},
			"4x²",
		},
		{
			poly.Int64P{
				poly.Int64T{poly.Ind{0, 1}, 1},
			},
			poly.Int64P{
				poly.Int64T{poly.Ind{1, 0}, 2},
			},
			"2xy",
		},
		{
			poly.Int64P{
				poly.Int64T{poly.Ind{0, 1}, 1},
				poly.Int64T{poly.Ind{1, 0}, 1},
				poly.Int64T{poly.Ind{1, 1}, 2},
			},
			poly.Int64P{
				poly.Int64T{poly.Ind{0, 0, 1}, 2},
			},
			"4xyz + 2xz + 2yz",
		},
		{
			poly.Int64P{
				poly.Int64T{poly.Ind{1, 0}, 1},
				poly.Int64T{poly.Ind{0, 1}, -1},
			},
			poly.Int64P{
				poly.Int64T{poly.Ind{1, 0}, 1},
				poly.Int64T{poly.Ind{0, 1}, -1},
			},
			"x² - 2xy + y²",
		},
		{
			poly.Int64P{
				poly.Int64T{poly.Ind{1, 0}, 1},
				poly.Int64T{poly.Ind{0, 1}, -1},
			},
			poly.Int64P{
				poly.Int64T{poly.Ind{1, 0}, 1},
				poly.Int64T{poly.Ind{0, 1}, 1},
			},
			"x² - y²",
		},
		{
			poly.Int64P{
				poly.Int64T{poly.Ind{1, 0}, 1},
				poly.Int64T{poly.Ind{0, 0}, -3},
			},
			poly.Int64P{
				poly.Int64T{poly.Ind{1, 0}, 1},
				poly.Int64T{poly.Ind{0, 0}, -3},
			},
			"x² - 6x + 9",
		},
	} {
		if got, want := row.a.Mul(row.b).String(), row.s; got != want {
			t.Errorf("(%q).Mul(%q) = %q; want %q", row.a, row.b, got, want)
		}
	}
}

func TestInt64PCompact(t *testing.T) {
	for _, row := range []struct {
		p poly.Int64P
		s string
	}{
		{poly.Int64P{
			poly.Int64T{poly.Ind{1}, -2},
			poly.Int64T{poly.Ind{6}, 0},
			poly.Int64T{poly.Ind{8}, 10},
			poly.Int64T{poly.Ind{0}, 5},
		}, "10x⁸ - 2x + 5"},
		{poly.Int64P{
			poly.Int64T{poly.Ind{1, 0}, 2},
			poly.Int64T{poly.Ind{1, 0}, 2},
			poly.Int64T{poly.Ind{1, 0}, -4},
			poly.Int64T{poly.Ind{0, 1}, 1},
			poly.Int64T{poly.Ind{0, 1}, 1},
			poly.Int64T{poly.Ind{0, 1}, 1},
		}, "3y"},
		{poly.Int64P{
			poly.Int64T{poly.Ind{0, 0, 1}, -1},
			poly.Int64T{poly.Ind{2, 0, 0}, 1},
			poly.Int64T{poly.Ind{0, 1}, 5},
			poly.Int64T{poly.Ind{1, 0, 0}, -8},
		}, "x² - 8x + 5y - z"},
	} {
		if got, want := row.p.Compact().String(), row.s; got != want {
			t.Errorf("(%#v).String() = %q; want %q", row.p, got, want)
		}
	}
}

func TestInt64PString(t *testing.T) {
	for _, row := range []struct {
		p poly.Int64P
		s string
	}{
		{poly.Int64P{}, "0"},
		{poly.Int64P{poly.Int64T{poly.Ind{}, 0}}, "0"},
		{poly.Int64P{poly.Int64T{poly.Ind{1, 2, 3}, 0}}, "0"},
		{poly.Int64P{poly.Int64T{poly.Ind{1, 2, 3}, 10}}, "10xy²z³"},
		{poly.Int64P{poly.Int64T{poly.Ind{-1, 0, 1}, -20}}, "-20x¯¹z"},
		{poly.Int64P{poly.Int64T{poly.Ind{}, 5}}, "5"},
		{poly.Int64P{poly.Int64T{poly.Ind{0, 0}, -8}}, "-8"},
	} {
		if got, want := row.p.String(), row.s; got != want {
			t.Errorf("(%#v).String() = %q; want %q", row.p, got, want)
		}
	}
}
