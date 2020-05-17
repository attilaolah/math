package knot

func Unknot() *Knot {
	return &Knot{&Arc{}}
}

func Trefoil() *Knot {
	return SimpleKnot(3)
}

func FigureEight() *Knot {
	return SimpleKnot(4)
}

// SimpleKnot creates a simple Knot.
// Each Arc crosses over the start of its previous Arc in linkage.
func SimpleKnot(size int) *Knot {
	arcs := make([]*Arc, size)
	crosses := make([]*Cross, size)

	for i := range arcs {
		arcs[i] = &Arc{}
		crosses[i] = &Cross{}
	}
	for i, a := range arcs {
		a.Start = crosses[i]
		a.Start.Out = a
		a.Stop = crosses[(i+1)%size]
		a.Stop.In = a
		crosses[(size+i-1)%size].Over = a
	}

	return &Knot{arcs[0]}
}
