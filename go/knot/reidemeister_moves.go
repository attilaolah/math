package knot

import "errors"

var (
	SlideError   = errors.New("knot: cannot slide arc: must be next to the cross")
	UntwistError = errors.New("knot: cannot untwist cross: arc must cross itself")
)

// Twist performs the first Reidemeister move.
// The resulting new cross is returned for convenience.
func Twist(a *Arc, h Handedness) *Cross {
	if a.Start == nil && a.Stop == nil {
		// Unknot. Create the first cross and return.
		c := Cross{In: a, Out: a, Over: a, Handedness: h}
		a.Start = &c
		a.Stop = &c

		return &c
	}

	c := Cross{In: a, Handedness: h}
	c.Out = &Arc{Start: &c, Stop: a.Stop}
	if h == Right {
		c.Over = c.Out
	} else {
		c.Over = a
	}
	a.Stop.In = c.Out
	a.Stop = &c

	return &c
}

// TwistLeft is a shortcut for calling Twist with left handedness.
func TwistLeft(a *Arc) *Cross { return Twist(a, Left) }

// TwistRight is a shortcut for calling Twist with right handedness.
func TwistRight(a *Arc) *Cross { return Twist(a, Right) }

// Untwist undoes the Twist() operation.
func Untwist(c *Cross) error {
	if c.Over == c.In {
		// Untwist Twist(a, over=true)
		c.Over.Stop = c.Out.Stop
	} else if c.Over == c.Out {
		// Untwist Twist(a, over=false)
		c.Over.Start = c.In.Start
	} else {
		return UntwistError
	}

	return nil
}

// Poke performs the second Reidemeister move.
// The first parameter, 'over', ends up going over in both crosses.
// Note that no validation is done on whether the poke is a possible move, i.e. whether there is another arc that
// separates the two args passed in as parameters. The resulting two new crosses are returned for convenience.
func Poke(over, under *Arc) (*Cross, *Cross) {
	c1 := Cross{Over: over, In: under}
	c2 := Cross{Over: over}
	c1.Out = &Arc{Start: &c1, Stop: &c2}
	c2.In = c1.Out
	c2.Out = &Arc{Start: &c2, Stop: under.Stop}
	under.Stop.In = c2.Out
	under.Stop = &c1

	return &c1, &c2
}

// Slide performs the third Reidemeister move.
// Repeating the same operation twice undoes the slide.
func Slide(a *Arc, c *Cross) error {
	if (a.Start.Over == c.Over) && (a.Stop.Over == c.In || a.Stop.Over == c.Out) {
		slideCrosses(c, a.Stop, a.Start)
	} else if (a.Start.Over == c.In || a.Start.Over == c.Out) && (a.Stop.Over == c.Over) {
		slideCrosses(c, a.Start, a.Stop)
	} else {
		return SlideError
	}

	return nil
}

// Slide using three crosses. Parameters are:
// c1: top + middle arc.
// c2: top + bottom arc.
// c3: middle + bottom arc.
func slideCrosses(c1, c2, c3 *Cross) {
	c3.Over = c1.Over
	if c2.Over == c1.In {
		c2.Over = c1.Out
	} else {
		c2.Over = c1.In
	}
}
