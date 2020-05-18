package knot

// Right and Left 'handedness' values:
const (
	Right Handedness = true
	Left  Handedness = false
)

// Handedness groups crosses into left- and right-handed.
type Handedness bool

// A Cross is where an arc goes over another arc.
type Cross struct {
	// A cross has one arc over, one going in and one coming out.
	Over, In, Out *Arc

	// Handedness can be right-handed (true) or left-handed (false).
	Handedness Handedness
}

// String represents handadness as "R" (right) or "L" (left).
func (h Handedness) String() string {
	if h {
		return "R"
	}

	return "L"
}

// Left returns the arc to the left of the cross.
func (c Cross) Left() *Arc {
	if c.Handedness == Left {
		return c.In
	}

	return c.Out
}

// Right returns the arc to the right of the cross.
func (c Cross) Right() *Arc {
	if c.Handedness == Right {
		return c.In
	}

	return c.Out
}

// Reverse swaps the in- and out-going arcs. Handedness is not changed.
func (c *Cross) reverse() {
	c.In, c.Out = c.Out, c.In
}
