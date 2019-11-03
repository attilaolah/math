package poly_test

import (
	"strings"
	"testing"

	"github.com/attilaolah/math/poly"
)

func TestInt64MDet(t *testing.T) {
	for _, row := range []struct {
		m poly.Int64M
		s string
	}{
		{poly.Int64M{
			[]poly.Int64P{
				{poly.Int64T{poly.Ind{1}, 1}},
			}, 1,
		}, "x"},
		{poly.Int64M{
			[]poly.Int64P{
				{poly.Int64T{poly.Ind{1, 0, 0, 0}, 1}},
				{poly.Int64T{poly.Ind{0, 1, 0, 0}, 1}},
				{poly.Int64T{poly.Ind{0, 0, 1, 0}, 1}},
				{poly.Int64T{poly.Ind{0, 0, 0, 1}, 1}},
			}, 2,
		}, "x₀x₃ - x₁x₂"},
		{poly.Int64M{
			[]poly.Int64P{
				{poly.Int64T{poly.Ind{1, 0, 0, 0, 0, 0, 0, 0, 0}, 1}},
				{poly.Int64T{poly.Ind{0, 1, 0, 0, 0, 0, 0, 0, 0}, 1}},
				{poly.Int64T{poly.Ind{0, 0, 1, 0, 0, 0, 0, 0, 0}, 1}},
				{poly.Int64T{poly.Ind{0, 0, 0, 1, 0, 0, 0, 0, 0}, 1}},
				{poly.Int64T{poly.Ind{0, 0, 0, 0, 1, 0, 0, 0, 0}, 1}},
				{poly.Int64T{poly.Ind{0, 0, 0, 0, 0, 1, 0, 0, 0}, 1}},
				{poly.Int64T{poly.Ind{0, 0, 0, 0, 0, 0, 1, 0, 0}, 1}},
				{poly.Int64T{poly.Ind{0, 0, 0, 0, 0, 0, 0, 1, 0}, 1}},
				{poly.Int64T{poly.Ind{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1}},
			}, 3,
		}, "x₀x₄x₈ - x₀x₅x₇ - x₁x₃x₈ + x₁x₅x₆ + x₂x₃x₇ - x₂x₄x₆"},
		{poly.Int64M{
			[]poly.Int64P{
				{poly.Int64T{poly.Ind{}, 1}},
				{poly.Int64T{poly.Ind{}, 2}},
				{poly.Int64T{poly.Ind{}, 3}},
				{poly.Int64T{poly.Ind{}, 4}},
				{poly.Int64T{poly.Ind{}, 5}},
				{poly.Int64T{poly.Ind{}, 6}},
				{poly.Int64T{poly.Ind{}, 7}},
				{poly.Int64T{poly.Ind{}, 8}},
				{poly.Int64T{poly.Ind{}, 9}},
			}, 3,
		}, "0"},
	} {
		if got, want := row.m.Det().String(), row.s; got != want {
			t.Errorf("(\n%s\n).Det() = %q; want: %q", row.m, got, want)
		}
	}
}

func TestInt64MString(t *testing.T) {
	for _, row := range []struct {
		m poly.Int64M
		s string
	}{
		{poly.Int64M{}, "[]"},
		{poly.Int64M{
			[]poly.Int64P{
				{poly.Int64T{poly.Ind{}, 0}},
				{poly.Int64T{poly.Ind{1, 0, 1}, 2}},
				{poly.Int64T{poly.Ind{}, -6}},
			}, 3,
		}, "[0 2xz -6]"},
		{poly.Int64M{
			[]poly.Int64P{
				{poly.Int64T{poly.Ind{}, 0}},
				{poly.Int64T{poly.Ind{1, 0, 1}, 2}},
				{poly.Int64T{poly.Ind{}, -6}},
			}, 1,
		}, `
⎡  0⎤
⎢2xz⎥
⎣ -6⎦
`},
		{poly.Int64M{
			[]poly.Int64P{
				{poly.Int64T{poly.Ind{}, 0}},
				{poly.Int64T{poly.Ind{1, 0, 1}, 2}},
				{poly.Int64T{poly.Ind{}, -6}},
				{poly.Int64T{poly.Ind{}, -1}},
				{poly.Int64T{poly.Ind{1, 0, 1}, 2}, poly.Int64T{poly.Ind{0, 1, 1}, 1}},
				{poly.Int64T{poly.Ind{}, 8}},
			}, 3,
		}, `
⎡ 0        2xz -6⎤
⎣-1 (2xz + yz)  8⎦
`},
		{poly.Int64M{
			[]poly.Int64P{
				{poly.Int64T{poly.Ind{}, 0}},
				{poly.Int64T{poly.Ind{1, 0, 1}, 2}},
				{poly.Int64T{poly.Ind{}, -6}},
				{poly.Int64T{poly.Ind{}, -1}},
				{poly.Int64T{poly.Ind{1, 0, 1}, 2}, poly.Int64T{poly.Ind{0, 1, 1}, 1}},
				{poly.Int64T{poly.Ind{}, 8}},
			}, 2,
		}, `
⎡         0 2xz⎤
⎢        -6  -1⎥
⎣(2xz + yz)   8⎦
`},
		{poly.Int64M{
			[]poly.Int64P{
				{poly.Int64T{poly.Ind{1, 0, 0, 0}, 1}},
				{poly.Int64T{poly.Ind{0, 1, 0, 0}, 1}},
				{poly.Int64T{poly.Ind{0, 0, 1, 0}, 1}},
				{poly.Int64T{poly.Ind{0, 0, 0, 1}, 1}},
			}, 2,
		}, `
⎡x₀ x₁⎤
⎣x₂ x₃⎦
`},
	} {
		row.s = strings.TrimSpace(row.s)
		if got, want := row.m.String(), row.s; got != want {
			t.Errorf("(%#v).String() =\n%s\n want:\n%s", row.m, got, want)
		}
	}
}
