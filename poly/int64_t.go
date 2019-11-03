// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package poly

import (
	"fmt"
	"strings"
)

// Int64T is a single term containing an int64 coefficient.
type Int64T struct {
	Ind
	C int64
}

// Ind represents the indeterminates of a single term.
type Ind []int64

// Mul returns the product of 't' and 'x'.
func (t Int64T) Mul(x Int64T) Int64T {
	t.Ind = t.Ind.Mul(x.Ind)
	t.C *= x.C
	return t
}

// Less reports whether 't' should be sorted before 'x' in a polynomial.
func (t Int64T) Less(x Int64T) bool {
	for i, k := range t.Ind {
		if k == x.Ind[i] {
			continue
		}
		return k > x.Ind[i]
	}
	return false
}

// String returns a compact, human-readable representation of the term.
func (t Int64T) String() string {
	if t.C == 0 {
		return "0"
	}

	s := t.Ind.String()
	if t.C == 1 {
		return s
	}
	if s == "1" {
		return fmt.Sprintf("%d", t.C)
	}

	return fmt.Sprintf("%d%s", t.C, s)
}

// Mul returns the product of two indeterminates.
func (i Ind) Mul(x Ind) Ind {
	if len(i) < len(x) {
		i, x = x, i
	}

	ret := make(Ind, len(i))
	for k := range i {
		ret[k] = i[k]
		if k < len(x) {
			ret[k] += x[k]
		}
	}

	return ret
}

// Eq checks two indeterminates for equality.
func (i Ind) Eq(x Ind) bool {
	if len(i) != len(x) {
		return false
	}
	for i, k := range i {
		if k != x[i] {
			return false
		}
	}
	return true
}

// String returns a compact, human-readable representation of the indeterminates.
func (i Ind) String() string {
	var s []string
	const simple = "xyz"

	if size := len(i); size == 0 {
		return "1"
	} else if size <= len(simple) {
		s = strings.Split(simple, "")[:size]
	} else {
		for i := range i {
			s = append(s, "x"+sub(int64(i)))
		}
	}

	ret := ""
	for i, x := range i {
		if x != 0 {
			ret += s[i]
			if x != 1 {
				ret += sup(x)
			}
		}
	}
	if ret == "" {
		return "1"
	}

	return ret
}

var (
	sub10 = [11]rune{'₀', '₁', '₂', '₃', '₄', '₅', '₆', '₇', '₈', '₉', '₋'}
	sup10 = [11]rune{'⁰', '¹', '²', '³', '⁴', '⁵', '⁶', '⁷', '⁸', '⁹', '¯'}
)

func sub(i int64) string {
	if i == 0 {
		return string(sub10[0])
	}
	return smap(i, sub10)
}

func sup(i int64) string {
	return smap(i, sup10)
}

func smap(i int64, m [11]rune) string {
	s := ""
	neg := i < 0
	if neg {
		i = -i
	}
	for i > 0 {
		s = fmt.Sprintf("%c%s", m[i%10], s)
		i /= 10
	}
	if neg {
		s = fmt.Sprintf("%c%s", m[10], s)
	}
	return s
}
