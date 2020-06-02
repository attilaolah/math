package knot_test

import (
	"reflect"
	"testing"

	"github.com/attilaolah/math/go/knot"
)

func TestGrid(t *testing.T) {
	for i, row := range []struct {
		dir  knot.Directions
		grid knot.Grid
	}{
		{
			// Simple unknot.
			dir: knot.Directions{knot.TurnLeft, knot.TurnLeft, knot.TurnLeft, knot.TurnLeft},
			grid: knot.Grid{
				knot.Point{0, 0}:  knot.Cell{knot.E, knot.TurnLeft},
				knot.Point{0, 1}:  knot.Cell{knot.N, knot.TurnLeft},
				knot.Point{-1, 1}: knot.Cell{knot.W, knot.TurnLeft},
				knot.Point{-1, 0}: knot.Cell{knot.S, knot.TurnLeft},
			},
		},
		{
			// Twisted unknot.
			dir: knot.Directions{
				knot.Forward, knot.TurnLeft, knot.TurnLeft, knot.TurnLeft,
				knot.Under, knot.TurnRight, knot.TurnRight, knot.TurnRight,
			},
			grid: knot.Grid{
				knot.Point{0, 0}:   knot.Cell{knot.ES, knot.Forward},
				knot.Point{1, 0}:   knot.Cell{knot.E, knot.TurnLeft},
				knot.Point{1, 1}:   knot.Cell{knot.N, knot.TurnLeft},
				knot.Point{0, 1}:   knot.Cell{knot.W, knot.TurnLeft},
				knot.Point{0, -1}:  knot.Cell{knot.S, knot.TurnRight},
				knot.Point{-1, -1}: knot.Cell{knot.W, knot.TurnRight},
				knot.Point{-1, 0}:  knot.Cell{knot.N, knot.TurnRight},
			},
		},
		{
			// Simple trefoil.
			dir: knot.Directions{
				knot.TurnLeft,
				knot.Forward,
				knot.TurnLeft,
				knot.Forward,
				knot.Forward,
				knot.TurnLeft,
				knot.TurnLeft,
				knot.Forward,
				knot.TurnLeft,
				knot.Forward,
				knot.TurnLeft,
				knot.TurnLeft,
				knot.Under,
				knot.Forward,
				knot.TurnLeft,
				knot.Forward,
			},
			grid: knot.Grid{
				knot.Point{0, 0}:  knot.Cell{knot.E, knot.TurnLeft},
				knot.Point{0, 1}:  knot.Cell{knot.N, knot.Forward},
				knot.Point{0, 2}:  knot.Cell{knot.N, knot.TurnLeft},
				knot.Point{-1, 2}: knot.Cell{knot.NW, knot.Forward},
				knot.Point{-2, 2}: knot.Cell{knot.WS, knot.Forward},
				knot.Point{-3, 2}: knot.Cell{knot.W, knot.TurnLeft},
				knot.Point{-3, 1}: knot.Cell{knot.S, knot.TurnLeft},
				knot.Point{-2, 1}: knot.Cell{knot.SE, knot.Forward},
				knot.Point{-1, 1}: knot.Cell{knot.E, knot.TurnLeft},
				knot.Point{-1, 3}: knot.Cell{knot.N, knot.TurnLeft},
				knot.Point{-2, 3}: knot.Cell{knot.W, knot.TurnLeft},
				knot.Point{-2, 0}: knot.Cell{knot.S, knot.TurnLeft},
				knot.Point{-1, 0}: knot.Cell{knot.E, knot.Forward},
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
