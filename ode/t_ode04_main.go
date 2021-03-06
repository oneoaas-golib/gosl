// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"math"
	"testing"
	"time"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/la"
	"github.com/cpmech/gosl/mpi"
	"github.com/cpmech/gosl/ode"
	"github.com/cpmech/gosl/plt"
)

func main() {

	mpi.Start(false)
	defer func() {
		mpi.Stop(false)
	}()

	if mpi.Rank() == 0 {
		chk.PrintTitle("ode04: Hairer-Wanner VII-p376 Transistor Amplifier\n")
	}
	if mpi.Size() != 3 {
		chk.Panic(">> error: this test requires 3 MPI processors\n")
		return
	}

	// data
	UE, UB, UF, ALPHA, BETA := 0.1, 6.0, 0.026, 0.99, 1.0e-6
	R0, R1, R2, R3, R4, R5 := 1000.0, 9000.0, 9000.0, 9000.0, 9000.0, 9000.0
	R6, R7, R8, R9 := 9000.0, 9000.0, 9000.0, 9000.0
	W := 2.0 * 3.141592654 * 100.0

	// initial values
	xa := 0.0
	ya := []float64{0.0,
		UB,
		UB / (R6/R5 + 1.0),
		UB / (R6/R5 + 1.0),
		UB,
		UB / (R2/R1 + 1.0),
		UB / (R2/R1 + 1.0),
		0.0}

	// endpoint of integration
	xb := 0.05
	//xb = 0.0123 // OK
	//xb = 0.01235 // !OK

	// right-hand side of the amplifier problem
	w := make([]float64, 8) // workspace
	fcn := func(f []float64, dx, x float64, y []float64) error {
		UET := UE * math.Sin(W*x)
		FAC1 := BETA * (math.Exp((y[3]-y[2])/UF) - 1.0)
		FAC2 := BETA * (math.Exp((y[6]-y[5])/UF) - 1.0)
		la.VecFill(f, 0)
		switch mpi.Rank() {
		case 0:
			f[0] = y[0] / R9
		case 1:
			f[1] = (y[1]-UB)/R8 + ALPHA*FAC1
			f[2] = y[2]/R7 - FAC1
		case 2:
			f[3] = y[3]/R5 + (y[3]-UB)/R6 + (1.0-ALPHA)*FAC1
			f[4] = (y[4]-UB)/R4 + ALPHA*FAC2
			f[5] = y[5]/R3 - FAC2
			f[6] = y[6]/R1 + (y[6]-UB)/R2 + (1.0-ALPHA)*FAC2
			f[7] = (y[7] - UET) / R0
		}
		mpi.AllReduceSum(f, w)
		return nil
	}

	// Jacobian of the amplifier problem
	jac := func(dfdy *la.Triplet, dx, x float64, y []float64) error {
		FAC14 := BETA * math.Exp((y[3]-y[2])/UF) / UF
		FAC27 := BETA * math.Exp((y[6]-y[5])/UF) / UF
		if dfdy.Max() == 0 {
			dfdy.Init(8, 8, 16)
		}
		NU := 2
		dfdy.Start()
		switch mpi.Rank() {
		case 0:
			dfdy.Put(2+0-NU, 0, 1.0/R9)
			dfdy.Put(2+1-NU, 1, 1.0/R8)
			dfdy.Put(1+2-NU, 2, -ALPHA*FAC14)
			dfdy.Put(0+3-NU, 3, ALPHA*FAC14)
			dfdy.Put(2+2-NU, 2, 1.0/R7+FAC14)
		case 1:
			dfdy.Put(1+3-NU, 3, -FAC14)
			dfdy.Put(2+3-NU, 3, 1.0/R5+1.0/R6+(1.0-ALPHA)*FAC14)
			dfdy.Put(3+2-NU, 2, -(1.0-ALPHA)*FAC14)
			dfdy.Put(2+4-NU, 4, 1.0/R4)
			dfdy.Put(1+5-NU, 5, -ALPHA*FAC27)
		case 2:
			dfdy.Put(0+6-NU, 6, ALPHA*FAC27)
			dfdy.Put(2+5-NU, 5, 1.0/R3+FAC27)
			dfdy.Put(1+6-NU, 6, -FAC27)
			dfdy.Put(2+6-NU, 6, 1.0/R1+1.0/R2+(1.0-ALPHA)*FAC27)
			dfdy.Put(3+5-NU, 5, -(1.0-ALPHA)*FAC27)
			dfdy.Put(2+7-NU, 7, 1.0/R0)
		}
		return nil
	}

	// matrix "M"
	c1, c2, c3, c4, c5 := 1.0e-6, 2.0e-6, 3.0e-6, 4.0e-6, 5.0e-6
	var M la.Triplet
	M.Init(8, 8, 14)
	M.Start()
	NU := 1
	switch mpi.Rank() {
	case 0:
		M.Put(1+0-NU, 0, -c5)
		M.Put(0+1-NU, 1, c5)
		M.Put(2+0-NU, 0, c5)
		M.Put(1+1-NU, 1, -c5)
		M.Put(1+2-NU, 2, -c4)
		M.Put(1+3-NU, 3, -c3)
	case 1:
		M.Put(0+4-NU, 4, c3)
		M.Put(2+3-NU, 3, c3)
		M.Put(1+4-NU, 4, -c3)
	case 2:
		M.Put(1+5-NU, 5, -c2)
		M.Put(1+6-NU, 6, -c1)
		M.Put(0+7-NU, 7, c1)
		M.Put(2+6-NU, 6, c1)
		M.Put(1+7-NU, 7, -c1)
	}

	// flags
	fixstp := false
	//method := "Dopri5"
	method := "Radau5"
	ndim := len(ya)
	numjac := false

	// ODE solver
	var o ode.Solver
	o.SaveXY = true
	o.Pll = true

	// solve problem
	if numjac {
		o.Init(method, ndim, fcn, nil, &M, nil)
	} else {
		o.Init(method, ndim, fcn, jac, &M, nil)
	}
	o.IniH = 1.0e-6 // initial step size

	// set tolerances
	atol, rtol := 1e-11, 1e-5
	o.SetTol(atol, rtol)

	// run
	t0 := time.Now()
	if fixstp {
		o.Solve(ya, xa, xb, 0.01, fixstp)
	} else {
		o.Solve(ya, xa, xb, xb-xa, fixstp)
	}

	// check
	if mpi.Rank() == 0 {
		tst := new(testing.T)
		chk.Int(tst, "number of F evaluations ", o.Nfeval, 2609)
		chk.Int(tst, "number of J evaluations ", o.Njeval, 215)
		chk.Int(tst, "total number of steps   ", o.Nsteps, 278)
		chk.Int(tst, "number of accepted steps", o.Naccepted, 221)
		chk.Int(tst, "number of rejected steps", o.Nrejected, 18)
		chk.Int(tst, "number of decompositions", o.Ndecomp, 276)
		chk.Int(tst, "number of lin solutions ", o.Nlinsol, 795)
		chk.Int(tst, "max number of iterations", o.Nitmax, 5)
	}

	// plot
	if mpi.Rank() == 0 {
		io.Pfmag("elapsed time = %v\n", time.Now().Sub(t0))
		plt.Reset(true, &plt.A{WidthPt: 450, Dpi: 150, Prop: 1.8, FszXtck: 6, FszYtck: 6})
		_, T, err := io.ReadTable("data/radau5_hwamplifier.dat")
		if err != nil {
			chk.Panic("%v", err)
		}
		s := o.IdxSave
		for j := 0; j < ndim; j++ {
			labelA, labelB := "", ""
			if j == 4 {
				labelA, labelB = "reference", "gosl"
			}
			plt.Subplot(ndim+1, 1, j+1)
			plt.Plot(T["x"], T[io.Sf("y%d", j)], &plt.A{C: "k", M: "+", L: labelA})
			plt.Plot(o.Xvalues[:s], o.Yvalues[j][:s], &plt.A{C: "r", M: ".", Ms: 1, Ls: "none", L: labelB})
			plt.AxisXmax(0.05)
			plt.Gll("$x$", io.Sf("$y_%d$", j), nil)
		}
		plt.Subplot(ndim+1, 1, ndim+1)
		plt.Plot(o.Xvalues[1:s], o.Hvalues[1:s], &plt.A{C: "b", NoClip: true})
		plt.SetYlog()
		plt.AxisXmax(0.05)
		plt.Gll("$x$", "$\\log{(h)}$", nil)
		plt.Save("/tmp/gosl", "hwamplifier_mpi")
	}
}
