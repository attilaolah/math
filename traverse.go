package math

// Traverse2d generates (x, y) pairs.
//
// This function iterates through points on a plane.
//
// At least one of (x, y) will be >= cur.
//
// If limit is nonzero, all coordinates will be <= limit. Note that the zero
// limit means "no limit", i.e. the channel will never be closed.
//
// Yields points in this order:
//
//    (0, 0) | cur = 0
//
//    (0, 1) | cur = 1
//    (1, 0)
//    (1, 1)
//
//    (0, 2) | cur = 2
//    (1, 2)
//    (2, 0)
//    (2, 1)
//    (2, 2)
//
//    (0, 3) | cur = 3
//    …
func Traverse2d(cur, limit uint64) <-chan [2]uint64 {
	ch := make(chan [2]uint64)
	go func() {
		for limit == 0 || cur <= limit {
			for x := uint64(0); x < cur; x++ {
				ch <- [2]uint64{x, cur}
			}
			for y := uint64(0); y < cur; y++ {
				ch <- [2]uint64{cur, y}
			}
			ch <- [2]uint64{cur, cur}
			cur++
		}
		close(ch)
	}()
	return ch
}

// Traverse3d Generates (x, y, z) triplets.
//
// This function iterates through points in 3D space.
//
// At least one of (x, y, z) will be >= cur.
//
// If limit is nonzero, all coordinates will be <= limit. Note that the zero
// limit means "no limit", i.e. the channel will never be closed.
//
// Yields points in this order:
//
//    (0, 0, 0) | cur = 0
//
//    (0, 0, 1) | cur = 1
//    (0, 1, 0)
//    (0, 1, 1)
//    (1, 0, 0)
//    (1, 0, 1)
//    (1, 1, 0)
//    (1, 1, 1)
//
//    (0, 0, 2) | cur = 2
//    (0, 1, 2)
//    (0, 2, 0)
//    (0, 2, 1)
//    (0, 2, 2)
//    (1, 0, 2)
//    (1, 1, 2)
//    (1, 2, 0)
//    (1, 2, 1)
//    (1, 2, 2)
//    (2, 0, 0)
//    (2, 0, 1)
//    (2, 1, 0)
//    (2, 1, 1)
//    (2, 0, 2)
//    (2, 1, 2)
//    (2, 2, 0)
//    (2, 2, 1)
//    (2, 2, 2)
//
//    (0, 0, 3) | cur = 3
//    …
func Traverse3d(cur, limit uint64) <-chan [3]uint64 {
	ch := make(chan [3]uint64)
	go func() {
		for limit == 0 || cur <= limit {
			for x := uint64(0); x < cur; x++ {
				for yz := range Traverse2d(cur, cur) {
					ch <- [3]uint64{x, yz[0], yz[1]}
				}
			}
			if cur != 0 {
				for yz := range Traverse2d(0, cur) {
					ch <- [3]uint64{cur, yz[0], yz[1]}
				}
			}
			cur++
		}
		close(ch)
	}()
	return ch
}
