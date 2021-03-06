// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package la

import (
	"math"
	"testing"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
)

func TestVector01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("TestVector 01")

	io.Pfyel("func VecFill(v []float64, s float64)\n")
	v := make([]float64, 5)
	VecFill(v, 666)
	PrintVec("v", v, "%5g", false)
	chk.Vector(tst, "v", 1e-17, v, []float64{666, 666, 666, 666, 666})

	io.Pfyel("\nfunc VecFillC(v []complex128, s complex128)\n")
	vc := make([]complex128, 5)
	VecFillC(vc, 666+666i)
	PrintVecC("vc", vc, "(%2g +", "%4gi) ", false)
	chk.VectorC(tst, "vc", 1e-17, vc, []complex128{666 + 666i, 666 + 666i, 666 + 666i, 666 + 666i, 666 + 666i})

	io.Pfyel("func VecApplyFunc(v []float64, f func(i int, x float64) float64)\n")
	vf := []float64{123, 120, 666}
	VecApplyFunc(vf, func(i int, x float64) float64 { return float64(i+1) + x/3 })
	PrintVec("vf", vf, "%5g", false)
	chk.Vector(tst, "vf", 1e-17, vf, []float64{42, 42, 225})

	io.Pfyel("func VecGetMapped(v []float64, f func(i int) float64) (v []float64)\n")
	vg := VecGetMapped(3, func(i int) float64 { return float64(i + 1) })
	PrintVec("vg", vg, "%5g", false)
	chk.Vector(tst, "vg", 1e-17, vg, []float64{1, 2, 3})

	io.Pfyel("func VecClone(a []float64) (b []float64)\n")
	va := []float64{1, 2, 3, 4, 5, 6}
	vb := VecClone(va)
	PrintVec("vb", vb, "%5g", false)
	chk.Vector(tst, "vb==va", 1e-17, vb, va)

	io.Pfyel("\nfunc VecAccum(v []float64) (sum float64)\n")
	PrintVec("v", v, "%5g", false)
	sum := VecAccum(v)
	io.Pf("sum(v) = %23.15e\n", sum)
	chk.Scalar(tst, "sum(v)", 1e-17, sum, 5*666)

	io.Pfyel("\nfunc VecNorm(v []float64) (nrm float64)\n")
	PrintVec("v", v, "%5g", false)
	nrm := VecNorm(v)
	io.Pf("norm(v) = %23.15e\n", nrm)
	chk.Scalar(tst, "norm(v)", 1e-17, nrm, 1.489221273014860e+03)

	io.Pfyel("\nfunc VecNormDiff(u, v []float64) (nrm float64)\n")
	u := []float64{333, 333, 333, 333, 333}
	PrintVec("u", u, "%5g", false)
	PrintVec("v", v, "%5g", false)
	nrm = VecNormDiff(u, v)
	io.Pf("norm(u-v) = %23.15e\n", nrm)
	chk.Scalar(tst, "norm(u-v)", 1e-17, nrm, math.Sqrt(5.0*333.0*333.0))

	io.Pfyel("\nfunc VecDot(u, v []float64) (res float64)\n")
	u = []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	PrintVec("u", u, "%5g", false)
	PrintVec("v", v, "%5g", false)
	udotv := VecDot(u, v)
	io.Pf("u dot v = %v\n", udotv)
	chk.Scalar(tst, "u dot v", 1e-12, udotv, 999)

	io.Pfyel("\nfunc VecCopy(a []float64, alp float64, b []float64)\n")
	a := make([]float64, len(u))
	VecCopy(a, 1, u)
	PrintVec("u     ", u, "%5g", false)
	PrintVec("a := u", a, "%5g", false)
	chk.Vector(tst, "a", 1e-17, a, []float64{0.1, 0.2, 0.3, 0.4, 0.5})

	io.Pfyel("\nfunc VecAdd(a []float64, alp float64, b []float64)\n")
	b := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	PrintVec("b        ", b, "%5g", false)
	VecAdd(b, 10, b) // b += 10.0*b
	PrintVec("b += 10*b", b, "%5g", false)
	chk.Vector(tst, "b", 1e-17, b, []float64{11, 22, 33, 44, 55})

	io.Pfyel("\nfunc VecAdd2(u []float64, alp float64, a []float64, bet float64, b []float64)\n")
	PrintVec("a", a, "%7g", false)
	PrintVec("b", b, "%7g", false)
	c := make([]float64, len(a))
	VecAdd2(c, 1, a, 10, b) // c = 1.0*a + 10.0*b
	PrintVec("c = 1*a+10*b", c, "%7g", false)
	chk.Vector(tst, "c", 1e-17, c, []float64{110.1, 220.2, 330.3, 440.4, 550.5})

	io.Pfyel("\nfunc VecMin(v []float64) (min float64)\n")
	PrintVec("a", a, "%5g", false)
	mina := VecMin(a)
	io.Pf("min(a) = %v\n", mina)
	chk.Scalar(tst, "min(a)", 1e-17, mina, 0.1)

	io.Pfyel("\nfunc VecMax(v []float64) (max float64)\n")
	PrintVec("a", a, "%5g", false)
	maxa := VecMax(a)
	io.Pf("max(a) = %v\n", maxa)
	chk.Scalar(tst, "max(a)", 1e-17, maxa, 0.5)

	io.Pfyel("\nfunc VecMinMax(v []float64) (min, max float64)\n")
	PrintVec("a", a, "%5g", false)
	min2a, max2a := VecMinMax(a)
	io.Pf("min(a) = %v\n", min2a)
	io.Pf("max(a) = %v\n", max2a)
	chk.Scalar(tst, "min(a)", 1e-17, min2a, 0.1)
	chk.Scalar(tst, "max(a)", 1e-17, max2a, 0.5)

	io.Pfyel("\nfunc VecLargest(u []float64, den float64) (largest float64)\n")
	PrintVec("b     ", b, "%5g", false)
	bdiv11 := []float64{b[0] / 11.0, b[1] / 11.0, b[2] / 11.0, b[3] / 11.0, b[4] / 11.0}
	PrintVec("b / 11", bdiv11, "%5g", false)
	maxbdiv11 := VecLargest(b, 11)
	io.Pf("max(b/11) = %v\n", maxbdiv11)
	chk.Scalar(tst, "max(b/11)", 1e-17, maxbdiv11, 5)

	io.Pfyel("\nfunc VecMaxDiff(a, b []float64) (maxdiff float64)\n")
	amb1 := []float64{a[0] - b[0], a[1] - b[1], a[2] - b[2], a[3] - b[3], a[4] - b[4]}
	amb2 := make([]float64, len(a))
	VecAdd2(amb2, 1, a, -1, b)
	PrintVec("a  ", a, "%7g", false)
	PrintVec("b  ", b, "%7g", false)
	PrintVec("a-b", amb1, "%7g", false)
	PrintVec("a-b", amb2, "%7g", false)
	maxdiffab := VecMaxDiff(a, b)
	io.Pf("maxdiff(a,b) = max(abs(a-b)) = %v\n", maxdiffab)
	chk.Vector(tst, "amb1 == amb2", 1e-17, amb1, amb2)
	chk.Scalar(tst, "maxdiff(a,b)", 1e-17, maxdiffab, 54.5)

	io.Pfyel("\nfunc VecMaxDiffC(a, b []complex128) (maxdiff float64)\n")
	az := []complex128{complex(a[0], 1), complex(a[1], 3), complex(a[2], 0.5), complex(a[3], 1), complex(a[4], 0)}
	bz := []complex128{complex(b[0], 1), complex(b[1], 6), complex(b[2], 0.8), complex(b[3], -3), complex(b[4], 1)}
	ambz := []complex128{az[0] - bz[0], az[1] - bz[1], az[2] - bz[2], az[3] - bz[3], az[4] - bz[4]}
	PrintVecC("az   ", az, "(%5g +", "%4gi) ", false)
	PrintVecC("bz   ", bz, "(%5g +", "%4gi) ", false)
	PrintVecC("az-bz", ambz, "(%5g +", "%4gi) ", false)
	maxdiffabz := VecMaxDiffC(az, bz)
	io.Pf("maxdiff(az,bz) = %v\n", maxdiffabz)
	chk.Scalar(tst, "maxdiff(az,bz)", 1e-17, maxdiffabz, 54.5)

	io.Pfyel("\nfunc VecScale(res []float64, Atol, Rtol float64, v []float64)\n")
	scal1 := make([]float64, len(a))
	VecScale(scal1, 0.5, 0.1, amb1)
	PrintVec("a-b            ", amb1, "%7g", false)
	PrintVec("0.5 + 0.1*(a-b)", scal1, "%7g", false)
	chk.Vector(tst, "0.5 + 0.1*(a-b)", 1e-15, scal1, []float64{-0.59, -1.68, -2.77, -3.86, -4.95})

	io.Pfyel("\nfunc VecScaleAbs(res []float64, Atol, Rtol float64, v []float64)\n")
	scal2 := make([]float64, len(a))
	VecScaleAbs(scal2, 0.5, 0.1, amb1)
	PrintVec("a-b            ", amb1, "%7g", false)
	PrintVec("0.5 + 0.1*|a-b|", scal2, "%7g", false)
	chk.Vector(tst, "0.5 + 0.1*|a-b|", 1e-15, scal2, []float64{1.59, 2.68, 3.77, 4.86, 5.95})

	io.Pfyel("\nfunc VecRms(u []float64) (rms float64)\n")
	PrintVec("v", v, "%5g", false)
	rms := VecRms(v)
	io.Pf("rms(v) = %23.15e\n", rms)
	chk.Scalar(tst, "rms(v)", 1e-17, rms, 666.0)

	io.Pfyel("func VecRmsErr(u []float64, Atol, Rtol float64, v []float64) (rms float64)\n")
	PrintVec("v", v, "%5g", false)
	rmserr := VecRmsErr(v, 0, 1, v)
	io.Pf("rmserr(v,v) = %23.15e\n", rmserr)
	chk.Scalar(tst, "rmserr(v,v,0,1)", 1e-17, rmserr, 1)

	io.Pfyel("func VecRmsError(u, w []float64, Atol, Rtol float64, v []float64) (rms float64)\n")
	PrintVec("v", v, "%5g", false)
	w := []float64{333, 333, 333, 333, 333}
	rmserr = VecRmsError(v, w, 0, 1, v)
	io.Pf("rmserr(v,w,v) = %23.15e\n", rmserr)
	chk.Scalar(tst, "rmserr(v,w,0,1,v)", 1e-17, rmserr, 0.5)
}
