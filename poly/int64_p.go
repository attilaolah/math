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
	"sort"
	"strings"
)

// Int64P is a polynomial with int64 terms and coefficients.
// This type implements a sparse representation, i.e. only non-zero terms are stored.
type Int64P []Int64T

// Add calculates the sum of two polynomials.
func (p Int64P) Add(x Int64P) Int64P {
	return append(p, x...).Compact()
}

// Mul calculates the product of two polynomials.
func (p Int64P) Mul(x Int64P) Int64P {
	ret := Int64P{}

	for _, t := range x {
		ret = append(ret, p.MulT(t)...)
	}
	return ret.Compact()
}

// MulT calculates the result of multiplying the polynomial by a single-term polynomial.
func (p Int64P) MulT(t Int64T) Int64P {
	ret := make(Int64P, len(p))
	copy(ret, p)

	for i := range p {
		ret[i] = p[i].Mul(t)
	}

	return ret
}

// Compact merges terms with the same indeterminates.
func (p Int64P) Compact() Int64P {
	p.Sort()

	ret := Int64P{}
	for _, t := range p {
		size := len(ret)
		if size == 0 {
			ret = append(ret, t)
			continue
		}
		if ret[size-1].Ind.Eq(t.Ind) {
			ret[size-1].C += t.C
			continue
		}
		ret = append(ret, t)
	}
	return ret
}

// String returns a compact, human-readable representation of the polynomial.
func (p Int64P) String() string {
	terms := []string{}
	for _, t := range p {
		if t.C == 0 {
			// Exclude "+ 0" terms.
			continue
		}
		if t.C > 0 {
			terms = append(terms, "+")
		} else {
			terms = append(terms, "-")
			t.C *= -1
		}
		terms = append(terms, t.String())
	}
	if len(terms) == 0 {
		return "0"
	}
	if terms[0] == "-" {
		terms[1] = "-" + terms[1]
	}
	return strings.Join(terms[1:], " ")
}

// Sort sorts the terms of the polynom, highest-first.
func (p Int64P) Sort() {
	sort.Sort(p)
}

// Len implements sort.Interface.
func (p Int64P) Len() int {
	return len(p)
}

// Less implements sort.Interface.
func (p Int64P) Less(i, j int) bool {
	return p[i].Less(p[j])
}

// Swap implements sort.Interface.
func (p Int64P) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
