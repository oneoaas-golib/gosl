// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oblas

import (
	"strings"

	"github.com/cpmech/gosl/io"
)

// Matrix implements a column-major representation of a matrix by using a linear array that can be passed to Fortran code
//
//  NOTE: the functions related to Matrix do not check for the limits of indices and dimensions.
//        Panic may occur then.
//
//  Example:
//             _      _
//            |  0  3  |
//        M = |  1  4  |
//            |_ 2  5 _|(m x n)
//
//     data[i+j*m] = M[i][j]
//
type Matrix struct {
	m, n int       // dimensions
	data []float64 // data array. column-major => Fortran
}

// NewMatrix allocates a new Matrix from given slice.
// NOTE: make sure to have at least 1x1 item
func NewMatrix(a [][]float64) (o *Matrix) {
	o = new(Matrix)
	o.m, o.n = len(a), len(a[0])
	o.data = make([]float64, o.m*o.n)
	o.SetFromSlice(a)
	return
}

// NewMatrixMN allocates a new (empty) Matrix with given MN (row/col sizes)
func NewMatrixMN(m, n int) (o *Matrix) {
	o = new(Matrix)
	o.m, o.n = m, n
	o.data = make([]float64, m*n)
	return
}

// SetFromSlice sets matrix with data from a nested slice structure
func (o *Matrix) SetFromSlice(a [][]float64) {
	k := 0
	for j := 0; j < o.n; j++ {
		for i := 0; i < o.m; i++ {
			o.data[k] = a[i][j]
			k += 1
		}
	}
}

// M returns the number of rows
func (o *Matrix) M() int { return o.m }

// N returns the number of columns
func (o *Matrix) N() int { return o.n }

// Set sets value
func (o *Matrix) Set(i, j int, val float64) {
	o.data[i+j*o.m] = val // col-major
}

// Get gets value
func (o *Matrix) Get(i, j int) float64 {
	return o.data[i+j*o.m] // col-major
}

// GetMat returns nested slice representation
func (o *Matrix) GetSlice() (M [][]float64) {
	M = make([][]float64, o.m)
	for i := 0; i < o.m; i++ {
		M[i] = make([]float64, o.n)
		for j := 0; j < o.n; j++ {
			M[i][j] = o.data[i+j*o.m]
		}
	}
	return
}

// Print prints matrix (without commas or brackets)
func (o *Matrix) Print(nfmt string) (l string) {
	if nfmt == "" {
		nfmt = "%g "
	}
	for i := 0; i < o.m; i++ {
		if i > 0 {
			l += "\n"
		}
		for j := 0; j < o.n; j++ {
			l += io.Sf(nfmt, o.Get(i, j))
		}
	}
	return
}

// PrintGo prints matrix in Go format
func (o *Matrix) PrintGo(nfmt string) (l string) {
	if nfmt == "" {
		nfmt = "%10g"
	}
	l = "[][]float64{\n"
	for i := 0; i < o.m; i++ {
		l += "    {"
		for j := 0; j < o.n; j++ {
			if j > 0 {
				l += ","
			}
			l += io.Sf(nfmt, o.Get(i, j))
		}
		l += "},\n"
	}
	l += "}"
	return
}

// PrintPy prints matrix in Python format
func (o *Matrix) PrintPy(nfmt string) (l string) {
	if nfmt == "" {
		nfmt = "%10g"
	}
	l = "np.matrix([\n"
	for i := 0; i < o.m; i++ {
		l += "    ["
		for j := 0; j < o.n; j++ {
			if j > 0 {
				l += ","
			}
			l += io.Sf(nfmt, o.Get(i, j))
		}
		l += "],\n"
	}
	l += "], dtype=float)"
	return
}

// complex ///////////////////////////////////////////////////////////////////////////////////////

// MatrixC implements a column-major representation of a matrix of complex numbers by using a linear
// array that can be passed to Fortran code.
//
//  NOTE: the functions related to MatrixC do not check for the limits of indices and dimensions.
//        Panic may occur then.
//
//  Example:
//             _            _
//            |  0+0i  3+3i  |
//        M = |  1+1i  4+4i  |
//            |_ 2+2i  5+5i _|(m x n)
//
//     data[i+j*m] = M[i][j]
//
type MatrixC struct {
	m, n int          // dimensions
	data []complex128 // data array. column-major => Fortran
}

// NewMatrixC allocates a new MatrixC from given slice.
// NOTE: make sure to have at least 1x1 items
func NewMatrixC(a [][]complex128) (o *MatrixC) {
	o = new(MatrixC)
	o.m, o.n = len(a), len(a[0])
	o.data = make([]complex128, o.m*o.n)
	o.SetFromSlice(a)
	return
}

// NewMatrixCmn allocates a new (empty) MatrixC with given mn (row/col sizes)
func NewMatrixCmn(m, n int) (o *MatrixC) {
	o = new(MatrixC)
	o.m, o.n = m, n
	o.data = make([]complex128, m*n)
	return
}

// SetFromSlice sets matrix with data from a nested slice structure
func (o *MatrixC) SetFromSlice(a [][]complex128) {
	k := 0
	for j := 0; j < o.n; j++ {
		for i := 0; i < o.m; i++ {
			o.data[k] = a[i][j]
			k += 1
		}
	}
}

// M returns the number of rows
func (o *MatrixC) M() int { return o.m }

// N returns the number of columns
func (o *MatrixC) N() int { return o.n }

// Set sets value
func (o *MatrixC) Set(i, j int, val complex128) {
	o.data[i+j*o.m] = val // col-major
}

// Get gets value
func (o *MatrixC) Get(i, j int) complex128 {
	return o.data[i+j*o.m] // col-major
}

// GetMat returns nested slice representation
func (o *MatrixC) GetSlice() (M [][]complex128) {
	M = make([][]complex128, o.m)
	for i := 0; i < o.m; i++ {
		M[i] = make([]complex128, o.n)
		for j := 0; j < o.n; j++ {
			M[i][j] = o.data[i+j*o.m]
		}
	}
	return
}

// Print prints matrix (without commas or brackets).
// NOTE: if non-empty, nfmtI must have '+' e.g. %+g
func (o *MatrixC) Print(nfmtR, nfmtI string) (l string) {
	if nfmtR == "" {
		nfmtR = "%g"
	}
	if nfmtI == "" {
		nfmtI = "%+g"
	}
	if !strings.ContainsRune(nfmtI, '+') {
		nfmtI = strings.Replace(nfmtI, "%", "%+", -1)
	}
	for i := 0; i < o.m; i++ {
		if i > 0 {
			l += "\n"
		}
		for j := 0; j < o.n; j++ {
			if j > 0 {
				l += ", "
			}
			l += io.Sf(nfmtR, real(o.Get(i, j))) + io.Sf(nfmtI, imag(o.Get(i, j))) + "i"
		}
	}
	return
}

// PrintGo prints matrix in Go format
// NOTE: if non-empty, nfmtI must have '+' e.g. %+g
func (o *MatrixC) PrintGo(nfmtR, nfmtI string) (l string) {
	if nfmtR == "" {
		nfmtR = "%g"
	}
	if nfmtI == "" {
		nfmtI = "%+g"
	}
	if !strings.ContainsRune(nfmtI, '+') {
		nfmtI = strings.Replace(nfmtI, "%", "%+", -1)
	}
	l = "[][]complex128{\n"
	for i := 0; i < o.m; i++ {
		l += "    {"
		for j := 0; j < o.n; j++ {
			if j > 0 {
				l += ","
			}
			l += io.Sf(nfmtR, real(o.Get(i, j))) + io.Sf(nfmtI, imag(o.Get(i, j))) + "i"
		}
		l += "},\n"
	}
	l += "}"
	return
}

// PrintPy prints matrix in Python format
// NOTE: if non-empty, nfmtI must have '+' e.g. %+g
func (o *MatrixC) PrintPy(nfmtR, nfmtI string) (l string) {
	if nfmtR == "" {
		nfmtR = "%g"
	}
	if nfmtI == "" {
		nfmtI = "%+g"
	}
	if !strings.ContainsRune(nfmtI, '+') {
		nfmtI = strings.Replace(nfmtI, "%", "%+", -1)
	}
	l = "np.matrix([\n"
	for i := 0; i < o.m; i++ {
		l += "    ["
		for j := 0; j < o.n; j++ {
			if j > 0 {
				l += ","
			}
			l += io.Sf(nfmtR, real(o.Get(i, j))) + io.Sf(nfmtI, imag(o.Get(i, j))) + "j"
		}
		l += "],\n"
	}
	l += "], dtype=complex)"
	return
}
