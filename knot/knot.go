package knot

import (
	"fmt"
	"strings"
)

// Knot represents a directed 2D knot diagram.
type Knot struct {
	// An arbitrary starting arc.
	start *Arc
}

// Arcs returns the arcs in order of linkage, by arc direction.
func (k Knot) Arcs() []*Arc {
	a := k.start
	arcs := []*Arc{a}

	for a = a.Next(); a != k.start && a != nil; a = a.Next() {
		arcs = append(arcs, a)
	}

	return arcs
}

// Crosses returns crosses corresponding to each arc, in order.
func (k Knot) Crosses() []*Cross {
	if k.start.Start == nil {
		// Unknot. There are no crosses.
		return nil
	}

	arcs := k.Arcs()
	crosses := make([]*Cross, len(arcs))
	for i, a := range arcs {
		crosses[i] = a.Start
	}

	return crosses
}

// Size returns the number of crosses in the knot.
func (k Knot) Size() int {
	return len(k.Crosses())
}

// Reverse changes the directionality of the knot.
// Note that the handedness of crosses stays the same.
func (k *Knot) Reverse() {
	crosses := k.Crosses()
	if crosses == nil {
		return
	}

	for _, c := range crosses {
		c.reverse()
		c.In.reverse()
	}

	k.start = crosses[0].Out
}

// String returns a visual representation of the knot.
func (k Knot) String() string {
	parts := []string{}
	arcs, crosses := k.Arcs(), k.Crosses()

	for i, a := range arcs {
		if a.Start != nil {
			parts = append(parts, fmt.Sprintf("%s%d", a.Start.Handedness, i+1))
		}
		s := fmt.Sprintf("A%d", i+1)
		over := []string{}
		for j, c := range crosses {
			if c.Over == a {
				over = append(over, fmt.Sprintf("%s%d", c.Handedness, j+1))
			}
		}
		if len(over) != 0 {
			s = fmt.Sprintf("%s{%s}", s, strings.Join(over, ", "))
		}
		parts = append(parts, s)
	}
	if a := arcs[len(arcs)-1]; a.Stop != nil {
		parts = append(parts, fmt.Sprintf("%s1", a.Stop.Handedness))
	}

	return strings.Join(parts, " ")
}
