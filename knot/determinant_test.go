package knot_test

import (
	"testing"

	"github.com/attilaolah/math/knot"
)

func TestDet(t *testing.T) {
	for i, row := range []struct {
		k   *knot.Knot
		det uint64
	}{
		{knot.Unknot(), 1},
		{knot.Trefoil(), 3},
		{knot.FigureEight(), 5},
		{knot.SimpleKnot(5), 11},
		{knot.SimpleKnot(6), 21},
		{knot.SimpleKnot(7), 43},
	} {
		if got, want := row.k.Det(), row.det; got != want {
			t.Errorf("#%d: Det() = %d; want: %d", i+1, got, want)
		}
	}
}
