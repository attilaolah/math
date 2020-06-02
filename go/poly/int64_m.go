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
	"math/rand"
	"strings"
)

// Int64M is a matrix with polynomial elements that have int64 terms and coefficients.
// The matrix itself is implemented as a dense representation, i.e. all elements are stored.
type Int64M struct {
	// Elements contains the actual elements, top-left to bottom-right, row by row.
	Elements []Int64P
	// Stride is the length of each row in the matrix.
	Stride uint
}

// NewInt64M creates a new zero-filled matrix wich each element set to the constant value 0.
func NewInt64M(rows, cols uint) *Int64M {
	m := Int64M{
		Elements: make([]Int64P, rows*cols),
		Stride:   cols,
	}
	for i := range m.Elements {
		m.Elements[i] = Int64P{Int64T{}}
	}
	return &m
}

// Det calculates the determinant of the square matrix.
func (m Int64M) Det() Int64P {
	if m.Stride*m.Stride != uint(len(m.Elements)) {
		panic("math error: determinant of non-square matrix")
	}

	switch m.Stride {
	case 0:
		panic("math error: determinant of empty matrix")
	case 1:
		return m.Elements[0]
	}

	ret, sign := Int64P{}, int64(1)
	for i := uint(0); i < m.Stride; i++ {
		p := m.Elements[i].MulT(Int64T{Ind{}, sign})
		ret = ret.Add(p.Mul(m.Minor(0, i).Det()))
		sign *= -1
	}
	return ret
}

// Minor returns a copy of 'm' with the 'i'-th row and 'j'-th column removed.
func (m Int64M) Minor(i, j uint) Int64M {
	ret := Int64M{Stride: m.Stride - 1}

	for k, p := range m.Elements {
		if uint(k)/m.Stride == i {
			// Remove the i-th row.
			continue
		}
		if uint(k)%m.Stride == j {
			// Remove the j-th column.
			continue
		}
		ret.Elements = append(ret.Elements, p)
	}

	return ret
}

// AnyMinor randomly picks a minor and returns it.
func (m Int64M) AnyMinor() Int64M {
	return m.Minor(uint(rand.Intn(int(m.Stride))), uint(rand.Intn(int(m.Stride))))
}

// String returns a compact, human-readable representation of the matrix.
func (m Int64M) String() string {
	if len(m.Elements) == 0 || m.Stride == 0 {
		return "[]"
	}

	parts := make([]string, len(m.Elements))
	sizes := make([]int, m.Stride)

	for i, e := range m.Elements {
		s := e.String()
		if len(e) > 1 {
			s = fmt.Sprintf("(%s)", s)
		}
		if size := len(strings.Split(s, "")); size > sizes[uint(i)%m.Stride] {
			sizes[uint(i)%m.Stride] = size
		}
		parts[i] = s
	}
	for i, s := range parts {
		parts[i] = fmt.Sprintf(fmt.Sprintf("%%%dv", sizes[uint(i)%m.Stride]), s)
	}

	rows := []string{}
	for len(parts) > 0 {
		rows = append(rows, strings.Join(parts[:m.Stride], " "))
		parts = parts[m.Stride:]
	}
	if len(rows) == 1 {
		return "[" + rows[0] + "]"
	}

	ret := []string{}
	for i, row := range rows {
		switch i {
		case 0:
			ret = append(ret, "⎡"+row+"⎤")
		case len(rows) - 1:
			ret = append(ret, "⎣"+row+"⎦")
		default:
			ret = append(ret, "⎢"+row+"⎥")
		}
	}

	return strings.Join(ret, "\n")
}
