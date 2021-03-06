// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fun

import (
	"math"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/plt"
	"github.com/cpmech/gosl/utl"
)

// Interpolation grid kinds
var (

	// UniformGridKind defines the uniform 1D grid kind
	UniformGridKind = io.NewEnum("Uniform", "fun.uniform", "U", "Uniform 1D grid")

	// ChebyGaussGridKind defines the Chebyshev-Gauss 1D grid kind
	ChebyGaussGridKind = io.NewEnum("ChebyGauss", "fun.chebygauss", "CG", "Chebyshev-Gauss 1D grid")
)

// LagrangeInterp implements Lagrange interpolators associated with a grid X
//
//   An interpolant I^X_N{f} (associated with a grid X; of degree N; with N+1 points)
//   is expressed in the Lagrange form as follows:
//
//                     N
//         X          ————             X
//        I {f}(x) =  \     f(x[i]) ⋅ ℓ (x)
//         N          /                i
//                    ————
//                    i = 0
//
//   where ℓ^X_i(x) is the i-th Lagrange cardinal polynomial associated with grid X and given by:
//
//                 N
//         N      ━━━━    x  -  X[j]
//        ℓ (x) = ┃  ┃  —————————————           0 ≤ i ≤ N
//         i      ┃  ┃   X[i] - X[j]
//               j = 0
//               j ≠ i
//
type LagrangeInterp struct {
	N int       // degree: N = len(X)-1
	X []float64 // grid points: len(X) = P+1; generated in [-1, 1]
}

// NewLagrangeInterp allocates a new LagrangeInterp
//   N        -- degree
//   gridType -- type of grid; e.g. uniform
//   NOTE: the grid will be generated in [-1, 1]
func NewLagrangeInterp(N int, gridType io.Enum) (o *LagrangeInterp, err error) {
	if N < 0 {
		return nil, chk.Err("N must be at least equal to 0. N=%d is invalid\n", N)
	}
	o = new(LagrangeInterp)
	o.N = N
	switch gridType {
	case UniformGridKind:
		o.X = utl.LinSpace(-1, 1, N+1)
	case ChebyGaussGridKind:
		o.X = make([]float64, N+1)
		h := math.Pi / float64(2*(N+1))
		for i := 0; i < N+1; i++ {
			o.X[i] = -math.Cos(h * float64(2*i+1))
		}
	default:
		return nil, chk.Err("cannot create grid type %q\n", gridType)
	}
	return
}

// W computes the generating (nodal) polynomial associated with grid X. The nodal polynomial is the
// unique polynomial of degree N+1 and leading coefficient whose zeros are the N+1 nodes of X.
//
//                 N
//         X      ━━━━
//        W (x) = ┃  ┃ (x - X[i])
//        N+1     ┃  ┃
//               i = 0
//
func (o *LagrangeInterp) W(x float64) (w float64) {
	w = 1
	for i := 0; i < o.N+1; i++ {
		w *= x - o.X[i]
	}
	return
}

// L computes the i-th Lagrange cardinal polynomial ℓ^X_i(x) associated with grid X
//
//                 N
//         X      ━━━━    x  -  X[j]
//        ℓ (x) = ┃  ┃  —————————————           0 ≤ i ≤ N
//         i      ┃  ┃   X[i] - X[j]
//               j = 0
//               j ≠ i
//
//   Input:
//      i -- index of X[i] point
//      x -- where to evaluate the polynomial
//   Output:
//      lix -- ℓ^X_i(x)
func (o *LagrangeInterp) L(i int, x float64) (lix float64) {
	lix = 1
	for j := 0; j < o.N+1; j++ {
		if i != j {
			lix *= (x - o.X[j]) / (o.X[i] - o.X[j])
		}
	}
	return
}

// I computes the interpolation I^X_N{f}(x) @ x
//
//                     N
//         X          ————             X
//        I {f}(x) =  \     f(x[i]) ⋅ ℓ (x)
//         N          /                i
//                    ————
//                    i = 0
//
func (o *LagrangeInterp) I(x float64, f Ss) (ix float64, err error) {
	for i := 0; i < o.N+1; i++ {
		fxi, e := f(o.X[i])
		if e != nil {
			return 0, e
		}
		ix += fxi * o.L(i, x)
	}
	return
}

// EstimateLebesgue estimates the Lebesgue constant by using 10000 stations along [-1,1]
func (o *LagrangeInterp) EstimateLebesgue() (ΛN float64) {
	nsta := 10000 // generate several points along [-1,1]
	for j := 0; j < nsta; j++ {
		x := -1.0 + 2.0*float64(j)/float64(nsta-1)
		sum := math.Abs(o.L(0, x))
		for i := 1; i < o.N+1; i++ {
			sum += math.Abs(o.L(i, x))
		}
		if sum > ΛN {
			ΛN = sum
		}
	}
	return
}

// EstimateMaxErr estimates the maximum error using 10000 stations along [-1,1]
// This function also returns the location (xloc) of the estimated max error
//   Computes:
//             maxerr = max(|f(x) - I{f}(x)|)
func (o *LagrangeInterp) EstimateMaxErr(f Ss) (maxerr, xloc float64) {
	nsta := 10000 // generate several points along [-1,1]
	xloc = -1
	for i := 0; i < nsta; i++ {
		x := -1.0 + 2.0*float64(i)/float64(nsta-1)
		fx, err := f(x)
		if err != nil {
			chk.Panic("f(x) failed:%v\n", err)
		}
		ix, err := o.I(x, f)
		if err != nil {
			chk.Panic("I(x) failed:%v\n", err)
		}
		e := math.Abs(fx - ix)
		if e > maxerr {
			maxerr = e
			xloc = x
		}
	}
	return
}

// PlotLagInterpL plots cardinal polynomials ℓ
func PlotLagInterpL(N int, gridType io.Enum) {
	xx := utl.LinSpace(-1, 1, 201)
	yy := make([]float64, len(xx))
	o, _ := NewLagrangeInterp(N, gridType)
	for n := 0; n < N+1; n++ {
		for k, x := range xx {
			yy[k] = o.L(n, x)
		}
		plt.Plot(xx, yy, &plt.A{NoClip: true})
	}
	Y := make([]float64, N+1)
	plt.Plot(o.X, Y, &plt.A{C: "k", Ls: "none", M: "o", Void: true, NoClip: true})
	plt.Gll("$x$", "$\\ell(x)$", nil)
	plt.Cross(0, 0, &plt.A{C: "grey"})
	plt.HideAllBorders()
}

// PlotLagInterpW plots nodal polynomial
func PlotLagInterpW(N int, gridType io.Enum) {
	npts := 201
	xx := utl.LinSpace(-1, 1, npts)
	yy := make([]float64, len(xx))
	o, _ := NewLagrangeInterp(N, gridType)
	for k, x := range xx {
		yy[k] = o.W(x)
	}
	Y := make([]float64, len(o.X))
	plt.Plot(o.X, Y, &plt.A{C: "k", Ls: "none", M: "o", Void: true, NoClip: true})
	plt.Plot(xx, yy, &plt.A{C: "b", Lw: 1, NoClip: true})
	plt.Gll("$x$", "$w(x)$", nil)
	plt.Cross(0, 0, &plt.A{C: "grey"})
	plt.HideAllBorders()
}

// PlotLagInterpI plots Lagrange interpolation I(x) function for many degrees Nvalues
func PlotLagInterpI(Nvalues []int, gridType io.Enum, f Ss) {
	npts := 201
	xx := utl.LinSpace(-1, 1, npts)
	yy := make([]float64, len(xx))
	for k, x := range xx {
		yy[k], _ = f(x)
	}
	iy := make([]float64, len(xx))
	plt.Plot(xx, yy, &plt.A{C: "k", Lw: 4, NoClip: true})
	for _, N := range Nvalues {
		p, _ := NewLagrangeInterp(N, gridType)
		for k, x := range xx {
			iy[k], _ = p.I(x, f)
		}
		E, xloc := p.EstimateMaxErr(f)
		plt.AxVline(xloc, &plt.A{C: "k", Ls: ":"})
		plt.Plot(xx, iy, &plt.A{L: io.Sf("$N=%d\\;E=%.3e$", N, E), NoClip: true})
	}
	plt.Cross(0, 0, &plt.A{C: "grey"})
	plt.Gll("$x$", "$f(x)\\quad I{f}(x)$", nil)
	plt.HideAllBorders()
}
