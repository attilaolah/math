package knot

import (
	"errors"
	"fmt"
)

// Orientation values: East, North, West, South.
const (
	E Orientation = iota
	N
	W
	S

	// Orientations used for crosses:
	EN // E over N
	NW // N over W
	WS // W over S
	SE // S over E
	ES // E over S
	NE // N over E
	WN // W over N
	SW // S over W

	// Used to clamp values.
	MaxOrientation
	MaxBaseOrientation = MaxOrientation / 3
)

const (
	GoForward Direction = iota
	GoLeft
	GoUnder
	GoRight

	// Used to clamp values.
	MaxDirection
)

var IncompleteDirections = errors.New("knot: incomplete directions")

// Orientation encodes an absolute directionality.
type Orientation byte

// Direction encodes a relative directionality.
// It is effectively the same as a orientation but relative to the current value.
type Direction Orientation

// Directions are the minimum information required to reconstruct a knot diagram.
type Directions []Direction

// Cell combines orientation with direction to be stored in the coordinate system.
type Cell struct {
	Orientation
	Direction
}

// Point encodes a point on the X/Y coordinate system.
type Point struct{ X, Y int }

// Grid is a sparse mapping of poinds (coordinates) to cells.
type Grid map[Point]Cell

// InvalidCrossing indicates an invalid crossing.
type InvalidCrossing Point

// Error implements the error interface.
func (err InvalidCrossing) Error() string {
	return fmt.Sprintf("knot: invalid crossing at position %s", Point(err))
}

// Grid decodes directions into a grid containing cells.
// It uses (0, 0) as the starting point, and east as the initial orientation.
func (ds Directions) Grid() (Grid, error) {
	g := Grid{}

	if len(ds) == 0 {
		return nil, IncompleteDirections
	}

	pos, o := Point{}, E
	for _, d := range ds {
		c := Cell{o, d}
		if old, ok := g[pos]; ok {
			// Position already in use, check whether we can make a crossing here.
			if d != GoForward && d != GoUnder {
				return nil, fmt.Errorf("%w: direction must be %s or %s, got %s", InvalidCrossing(pos), GoForward, GoUnder, d)
			}
			if old.IsCross() {
				// Technically this should never happen, but still.
				return nil, fmt.Errorf("%w: coordinates already contain crossing %s", InvalidCrossing(pos), old)
			}
			if !old.IsStraight() {
				return nil, fmt.Errorf("%w: can only cross at straight line, not %s", InvalidCrossing(pos), old)
			}
			if !old.IsPerpendicular(o) {
				return nil, fmt.Errorf("%w: existing cell %s is not perpendicular to current orientation %s", InvalidCrossing(pos), old, o)
			}
			g[pos] = Cell{old.Cross(o, d == GoForward), GoForward}
		} else {
			// Position not in use.
			if d == GoUnder {
				return nil, fmt.Errorf("%w: nothing to go under", InvalidCrossing(pos))
			}
			g[pos] = c
		}
		o = o.Turn(d)
		pos = pos.Step(o)
	}
	if (pos != Point{}) {
		return nil, IncompleteDirections
	}

	return g, nil
}

// Base clamps the orientation to its base orientation.
func (o Orientation) Base() Orientation {
	return o % MaxBaseOrientation
}

// Cross crosses two orientations, with the second one going either over (over = true) or under (over = false).
func (o Orientation) Cross(other Orientation, over bool) Orientation {
	if (o == N && other == E && over) || (o == E && other == N && !over) {
		return EN
	}
	if (o == W && other == N && over) || (o == N && other == W && !over) {
		return NW
	}
	if (o == S && other == W && over) || (o == W && other == S && !over) {
		return WS
	}
	if (o == E && other == S && over) || (o == S && other == E && !over) {
		return SE
	}
	if (o == S && other == E && over) || (o == E && other == S && !over) {
		return ES
	}
	if (o == E && other == N && over) || (o == N && other == E && !over) {
		return NE
	}
	if (o == N && other == W && over) || (o == W && other == N && !over) {
		return WN
	}
	if (o == W && other == S && over) || (o == S && other == W && !over) {
		return SW
	}
	panic(fmt.Sprintf("invalid cross: %s.Cross(%s, %v)", o, other, over))
}

// IsCross checks whether the orientation corresponds to a crossing.
func (o Orientation) IsCross() bool {
	return o != o.Base()
}

// Clamp returns a copy with all but the two least significant bits cleared.
func (o Orientation) Clamp() Orientation {
	return o % MaxOrientation
}

// Turn towards direction.
func (o Orientation) Turn(to Direction) Orientation {
	switch to.Clamp() {
	case GoForward, GoUnder:
		return o
	case GoLeft:
		return (o + 1).Base()
	case GoRight:
		return (o - 1).Base()
	}
	panic("knot: should not happen")
}

// IsPerpendicular checks whether two orientations are perpendicular.
func (o Orientation) IsPerpendicular(other Orientation) bool {
	return (o+other)%2 == 1
}

// String returns a short, human-readable representation.
func (o Orientation) String() string {
	switch o.Clamp() {
	case E:
		return "E"
	case N:
		return "N"
	case W:
		return "W"
	case S:
		return "S"
	case EN:
		return "EN"
	case NW:
		return "NW"
	case WS:
		return "WS"
	case SE:
		return "SE"
	case ES:
		return "ES"
	case NE:
		return "NE"
	case WN:
		return "WN"
	case SW:
		return "SW"
	}
	panic("knot: should not happen")
}

// Clamp returns a copy with all but the two least significant bits cleared.
func (d Direction) Clamp() Direction {
	return d % MaxDirection
}

// IsStraight returns true if the direction is forward or under (i.e. do not turn).
func (d Direction) IsStraight() bool {
	return d%2 == 0
}

// String returns a short, human-readable representation.
func (d Direction) String() string {
	switch d.Clamp() {
	case GoForward:
		return "F"
	case GoLeft:
		return "L"
	case GoUnder:
		return "U"
	case GoRight:
		return "R"
	}
	panic("knot: should not happen")
}

// String returns a short, human-readable representation.
func (c Cell) String() string {
	return fmt.Sprintf("%s (%s)", c.Orientation, c.Direction)
}

// Step returns the point one step towards the passed orientation.
func (p Point) Step(towards Orientation) Point {
	switch towards.Base() {
	case E:
		return Point{p.X + 1, p.Y}
	case N:
		return Point{p.X, p.Y + 1}
	case W:
		return Point{p.X - 1, p.Y}
	case S:
		return Point{p.X, p.Y - 1}
	}
	panic("knot: should not happen")
}

// String returns a short, human-readable representation.
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}
