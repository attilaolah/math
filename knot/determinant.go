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
	crosses := k.Crosses()
	if len(crosses) == 0 {
		return nil
	}

	m := poly.NewInt64M(uint(len(crosses)), uint(len(crosses)))
	for row, rc := range crosses {
		for col, cc := range crosses {
			var f int64
			if rc.In == cc.Out {
				f -= 1
			}
			if rc.Out == cc.Out {
				f -= 1
			}
			if rc.Over == cc.Out {
				f += 2
			}
			m.Elements[row*int(m.Stride)+col][0].C = f
		}
	}

	return m
}
