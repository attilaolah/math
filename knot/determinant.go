package knot

import (
	"fmt"

	"github.com/attilaolah/math/poly"
)

// Det calculates the Knot's determinant.
func (k *Knot) Det() uint64 {
	m := k.Matrix()
	if m == nil {
		return 1
	}

	det := k.Matrix().AnyMinor().Det()
	if len(det) != 1 {
		panic(fmt.Sprintf("knot: unexpected determinant %s", det))
	}
	c := det[0].C
	if c < 0 {
		return uint64(-c)
	}
	return uint64(c)
}

// Matrix generates the matrix for calculating determinant of the Knot.
func (k *Knot) Matrix() *poly.Int64M {
	arcs, crosses := k.Arcs(), k.Crosses()
	if len(crosses) == 0 {
		return nil
	}

	m := poly.Int64M{
		Elements: make([]poly.Int64P, len(crosses)*len(arcs)),
		Stride:   uint(len(crosses)),
	}
	// TODO: Only traverse crosses to reduce complexity!
	for i, c := range crosses {
		for j, a := range arcs {
			var f int64
			if c.In == a {
				f -= 1
			}
			if c.Out == a {
				f -= 1
			}
			if c.Over == a {
				f += 2
			}
			m.Elements[i*int(m.Stride)+j] = poly.Int64P{poly.Int64T{C: f}}
		}
	}

	return &m
}
