package knot_test

import (
	"testing"

	"github.com/attilaolah/math/knot"
)

func TestArcs(t *testing.T) {
	for i, row := range []struct {
		k    *knot.Knot
		size int
	}{
		{knot.Unknot(), 1},
		{knot.Trefoil(), 3},
	} {
		k := row.k

		// Sanity check: make sure lenght is correct.
		if got, want := len(k.Arcs()), row.size; got != want {
			t.Errorf("#%d: k.Arcs() size = %d; want %d", i+1, got, want)
		}

		if row.size == 1 {
			// The unknot is not linked, nothing else to check here.
			continue
		}

		for rev := 0; rev < 3; rev++ {
			// Make sure arcs are linked in order.
			arcs := k.Arcs()
			for j, a := range arcs {
				if got, want := a.Next(), arcs[(j+1)%len(arcs)]; got != want {
					t.Errorf("#%d (rev %d): k.Arcs()[%d].Next() = %p; want %p", i+1, rev, j, got, want)
				}
				if got, want := a.Prev(), arcs[(len(arcs)+j-1)%len(arcs)]; got != want {
					t.Errorf("#%d (rev %d): k.Arcs()[%d].Prev() = %p; want %p", i+1, rev, j, got, want)
				}
			}

			// Reverse the knot and re-run these tests to make sure the order doesn't change.
			k.Reverse()
		}
	}
}

func TestCrosses(t *testing.T) {
	for i, row := range []struct {
		k    *knot.Knot
		size int
	}{
		{knot.Unknot(), 0},
		{knot.Trefoil(), 3},
	} {
		k := row.k

		// Sanity check: make sure lenght is correct.
		if got, want := len(k.Crosses()), row.size; got != want {
			t.Errorf("#%d: k.Crosses() size = %d; want %d", i+1, got, want)
		}

		for rev := 0; rev < 3; rev++ {
			// Make sure crosses are linked in order.
			crosses := k.Crosses()
			for j, c := range crosses {
				if got, want := c.Out.Stop, crosses[(j+1)%len(crosses)]; got != want {
					t.Errorf("#%d (rev: %d): k.Crosses()[%d].Out.Stop = %p; want %p", i+1, rev, j, got, want)
				}
				if got, want := c.In.Start, crosses[(len(crosses)+j-1)%len(crosses)]; got != want {
					t.Errorf("#%d (rev: %d): k.Crosses()[%d].In.Start = %p; want %p", i+1, rev, j, got, want)
				}
			}

			// Reverse the knot and re-run these tests to make sure the order doesn't change.
			k.Reverse()
		}
	}
}

func TestString(t *testing.T) {
	type row struct {
		k    *knot.Knot
		want string
	}
	rows := []row{
		{knot.Unknot(), "A1"},
		{knot.Trefoil(), "L1 A1{L3} L2 A2{L1} L3 A3{L2} L1"},
		{knot.FigureEight(), "L1 A1{L4} L2 A2{L1} L3 A3{L2} L4 A4{L3} L1"},
	}
	{
		k := knot.Unknot()
		knot.TwistLeft(k.Arcs()[0])
		rows = append(rows, row{k, "L1 A1{L1} L1"})
	}
	{
		k := knot.Unknot()
		knot.TwistRight(k.Arcs()[0])
		rows = append(rows, row{k, "R1 A1{R1} R1"})
	}
	{
		k := knot.Unknot()
		knot.TwistLeft(k.Arcs()[0])
		knot.TwistLeft(k.Arcs()[0])
		rows = append(rows, row{k, "L1 A1{L1, L2} L2 A2 L1"})
	}
	{
		k := knot.Unknot()
		knot.TwistLeft(k.Arcs()[0])
		knot.TwistRight(k.Arcs()[0])
		rows = append(rows, row{k, "L1 A1{L1} R2 A2{R2} L1"})
	}
	{
		k := knot.Trefoil()
		knot.TwistLeft(k.Arcs()[0])
		// TODO: Should be A1{L4, L2} to keep ordering of crosses.
		rows = append(rows, row{k, "L1 A1{L2, L4} L2 A2 L3 A3{L1} L4 A4{L3} L1"})
	}

	for i, row := range rows {
		if got, want := row.k.String(), row.want; got != want {
			t.Errorf("#%d: k.String():\n  %s\nwant:\n  %s", i+1, got, want)
		}
	}
}
