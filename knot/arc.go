package knot

// An Arc is a segment of the knot.
type Arc struct {
	// An Arc starts at 'Start' and ends at 'Stop'.
	// As a special case, the unknot has no start and stop.
	Start, Stop *Cross
}

// Prev returns the previous arc in the knot.
func (a Arc) Prev() *Arc {
	if a.Start == nil {
		// Unknot. There's no previous arc.
		return &a
	}

	return a.Start.In
}

// Next returs the next arc in the knot.
func (a Arc) Next() *Arc {
	if a.Stop == nil {
		// Unknot. There's no next arc.
		return &a
	}

	return a.Stop.Out
}

// Touches determines whether two arcs touch.
// Two arcs touch if they share a cross where one of them goes over.
func (a Arc) Touches(other *Arc) bool {
	if a.Start == nil {
		// The Unknot does not touch itself, unless twisted.
		return false
	}

	return a.Start.Over == other || a.Stop.Over == other || other.Start.Over == &a || other.Stop.Over == &a
}

// Reverse changes the direction of the arc.
func (a *Arc) reverse() {
	a.Start, a.Stop = a.Stop, a.Start
}
