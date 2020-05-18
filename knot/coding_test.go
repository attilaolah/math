package knot_test

import (
	"reflect"
	"testing"

	"github.com/attilaolah/math/knot"
)

func TestGrid(t *testing.T) {
	for i, row := range []struct {
		dir  knot.Directions
		grid knot.Grid
	}{
		{
			dir: knot.Directions{knot.GoLeft, knot.GoLeft, knot.GoLeft, knot.GoLeft},
			grid: knot.Grid{
				knot.Point{0, 0}:  knot.Cell{knot.E, knot.GoLeft},
				knot.Point{0, 1}:  knot.Cell{knot.N, knot.GoLeft},
				knot.Point{-1, 1}: knot.Cell{knot.W, knot.GoLeft},
				knot.Point{-1, 0}: knot.Cell{knot.S, knot.GoLeft},
			},
		},
		{
			dir: knot.Directions{
				knot.GoForward, knot.GoLeft, knot.GoLeft, knot.GoLeft,
				knot.GoUnder, knot.GoRight, knot.GoRight, knot.GoRight,
			},
			grid: knot.Grid{
				knot.Point{0, 0}:   knot.Cell{knot.ES, knot.GoForward},
				knot.Point{1, 0}:   knot.Cell{knot.E, knot.GoLeft},
				knot.Point{1, 1}:   knot.Cell{knot.N, knot.GoLeft},
				knot.Point{0, 1}:   knot.Cell{knot.W, knot.GoLeft},
				knot.Point{0, -1}:  knot.Cell{knot.S, knot.GoRight},
				knot.Point{-1, -1}: knot.Cell{knot.W, knot.GoRight},
				knot.Point{-1, 0}:  knot.Cell{knot.N, knot.GoRight},
			},
		},
	} {
		if grid, err := row.dir.Grid(); err != nil {
			t.Errorf("%d: Grid() returned error: %v", i+1, err)
		} else if !reflect.DeepEqual(grid, row.grid) {
			t.Errorf("#%d:\nDirections: %s\nGrid: %v\nWant: %v", i+1, row.dir, grid, row.grid)
		}
	}
}
