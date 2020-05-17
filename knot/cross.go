package knot

// Right and Left 'handedness' values:
const (
	Right Direction = true
	Left  Direction = false
)

// Direction groups crosses into left- and right-handed.
type Direction bool

// A Cross is where an arc goes over another arc.
type Cross struct {
	// A cross has one arc over, one going in and one coming out.
	Over, In, Out *Arc

	// Direction can be right-handed (true) or left-handed (false).
	Dir Direction
}

// String represents handadness as "R" (right) or "L" (left).
func (d Direction) String() string {
	if d {
		return "R"
	}

	return "L"
}

// Left returns the arc to the left of the cross.
func (c Cross) Left() *Arc {
	if c.Dir == Left {
		return c.In
	}

	return c.Out
}

// Right returns the arc to the right of the cross.
func (c Cross) Right() *Arc {
	if c.Dir == Right {
		return c.In
	}

	return c.Out
}

// Reverse swaps the in- and out-going arcs. Handedness is not changed.
func (c *Cross) reverse() {
	c.In, c.Out = c.Out, c.In
}
