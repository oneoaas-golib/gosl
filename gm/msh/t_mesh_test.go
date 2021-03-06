// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package msh

import (
	"testing"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/plt"
)

func Test_singleq4(tst *testing.T) {

	//verbose()
	chk.PrintTitle("singleq4")

	// load mesh
	m, err := Read("data/singleq4square1x1.msh")
	if err != nil {
		tst.Errorf("Read failed:\n%v\n", err)
		return
	}

	// correct data
	nverts := 4
	ncells := 1
	allX := [][]float64{
		{0, 0},
		{1, 0},
		{1, 1},
		{0, 1},
	}
	allVtags := []int{-1, 0, -1, 0}
	allCtags := []int{-1}
	allParts := []int{0}
	allTypeKeys := []string{"qua4"}
	allTypeIndices := []int{TypeQua4}
	allV := [][]int{
		{0, 1, 2, 3},
	}
	allEtags := [][]int{
		{-10, -11, -12, -13},
	}

	// check input data
	checkinput(tst, m, nverts, ncells, allX, allVtags, allCtags, allParts, allTypeKeys, allTypeIndices, allV, allEtags, nil)

	// derived data
	ndim := 2
	xmin := []float64{0.0, 0.0}
	xmax := []float64{1.0, 1.0}
	allGndim := []int{2}
	allCoords := [][][]float64{
		[][]float64{
			{0.0, 0.0},
			{1.0, 0.0},
			{1.0, 1.0},
			{0.0, 1.0},
		},
	}

	// check derived data
	checkderived(tst, m, ndim, xmin, xmax, allGndim, allCoords)

	// get map of tags
	tm, err := m.GetTagMaps()
	if err != nil {
		tst.Errorf("GetTagMaps failed:\n%v\n", err)
		return
	}

	// correct data
	vtags := []int{-1}
	ctags := []int{-1}
	cparts := []int{0}
	etags := []int{-10, -11, -12, -13}
	ctypeinds := []int{TypeQua4}
	vtagsVids := [][]int{{0, 2}}
	ctagsCids := [][]int{{0}}
	cpartsCids := [][]int{{0}}
	ctypesCids := [][]int{{0}}
	etagsCids := [][]int{{0}, {0}, {0}, {0}} // not unique
	etagsLocEids := [][]int{{0}, {1}, {2}, {3}}
	etagsVids := [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 3}}

	// check maps
	checkmaps(tst, m, tm, vtags, ctags, cparts, etags, nil, ctypeinds, vtagsVids, ctagsCids, cpartsCids, ctypesCids, etagsVids, etagsCids, etagsLocEids, nil, nil, nil)

	// draw
	if chk.Verbose {
		args := NewArgs()
		args.WithIdsCells = true
		plt.Reset(true, nil)
		m.Draw(args)
		plt.HideAllBorders()
		plt.Save("/tmp/gosl/gm", "singleq4")
	}
}

func Test_mesh01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("mesh01")

	// load mesh
	m, err := Read("data/mesh01.msh")
	if err != nil {
		tst.Errorf("Read failed:\n%v\n", err)
		return
	}

	// correct data
	nverts := 11
	ncells := 6
	allX := [][]float64{
		{0.0, 0.0},
		{0.5, 0.0},
		{1.0, 0.0},
		{1.4, 0.25},
		{0.0, 0.5},
		{0.5, 0.5},
		{1.0, 0.5},
		{1.4, 0.75},
		{0.0, 1.0},
		{0.5, 1.0},
		{1.0, 1.0},
	}
	allVtags := []int{-1, -1, -1, -2, -3, -3, -3, -4, -5, -5, -5}
	allCtags := []int{-1, -1, -1, -2, -2, -2}
	allParts := []int{0, 1, 2, 0, 1, 2}
	allTypeKeys := []string{"qua4", "qua4", "tri3", "qua4", "qua4", "tri3"}
	allTypeIndices := []int{TypeQua4, TypeQua4, TypeTri3, TypeQua4, TypeQua4, TypeTri3}
	allV := [][]int{
		{0, 1, 5, 4},
		{1, 2, 6, 5},
		{2, 3, 6},
		{4, 5, 9, 8},
		{5, 6, 10, 9},
		{6, 7, 10},
	}
	allEtags := [][]int{
		{-10, 0, 0, -13},
		{-10, 0, 0, 0},
		{-11, -11, 0},
		{0, 0, -12, -13},
		{0, 0, -12, 0},
		{-11, -11, 0},
	}

	// check input data
	checkinput(tst, m, nverts, ncells, allX, allVtags, allCtags, allParts, allTypeKeys, allTypeIndices, allV, allEtags, nil)

	// derived data
	ndim := 2
	xmin := []float64{0.0, 0.0}
	xmax := []float64{1.4, 1.0}
	allGndim := []int{2, 2, 2, 2, 2, 2}
	allCoords := [][][]float64{
		[][]float64{
			{0.0, 0.0},
			{0.5, 0.0},
			{0.5, 0.5},
			{0.0, 0.5},
		},
		[][]float64{
			{0.5, 0.0},
			{1.0, 0.0},
			{1.0, 0.5},
			{0.5, 0.5},
		},
		[][]float64{
			{1.0, 0.0},
			{1.4, 0.25},
			{1.0, 0.5},
		},
		[][]float64{
			{0.0, 0.5},
			{0.5, 0.5},
			{0.5, 1.0},
			{0.0, 1.0},
		},
		[][]float64{
			{0.5, 0.5},
			{1.0, 0.5},
			{1.0, 1.0},
			{0.5, 1.0},
		},
		[][]float64{
			{1.0, 0.5},
			{1.4, 0.75},
			{1.0, 1.0},
		},
	}

	// check derived data
	checkderived(tst, m, ndim, xmin, xmax, allGndim, allCoords)

	// get map of tags
	tm, err := m.GetTagMaps()
	if err != nil {
		tst.Errorf("GetTagMaps failed:\n%v\n", err)
		return
	}

	// correct data
	vtags := []int{-1, -2, -3, -4, -5}
	ctags := []int{-1, -2}
	cparts := []int{0, 1, 2}
	etags := []int{-10, -11, -12, -13}
	ctypeinds := []int{TypeQua4, TypeTri3}
	vtagsVids := [][]int{{0, 1, 2}, {3}, {4, 5, 6}, {7}, {8, 9, 10}}
	ctagsCids := [][]int{{0, 1, 2}, {3, 4, 5}}
	cpartsCids := [][]int{{0, 3}, {1, 4}, {2, 5}}
	ctypesCids := [][]int{{0, 1, 3, 4}, {2, 5}}
	etagsCids := [][]int{{0, 1}, {2, 2, 5, 5}, {3, 4}, {0, 3}} // not unique
	etagsLocEids := [][]int{{0, 0}, {0, 1, 0, 1}, {2, 2}, {3, 3}}
	etagsVids := [][]int{{0, 1, 2}, {2, 3, 6, 7, 10}, {8, 9, 10}, {0, 4, 8}}

	// check maps
	checkmaps(tst, m, tm, vtags, ctags, cparts, etags, nil, ctypeinds, vtagsVids, ctagsCids, cpartsCids, ctypesCids, etagsVids, etagsCids, etagsLocEids, nil, nil, nil)

	// draw
	if chk.Verbose {
		args := NewArgs()
		args.WithEdges = true
		args.WithIdsCells = true
		plt.Reset(true, nil)
		m.Draw(args)
		plt.HideAllBorders()
		plt.Save("/tmp/gosl/gm", "mesh01")
	}
}

func Test_cubeandtet(tst *testing.T) {

	//verbose()
	chk.PrintTitle("cubeandtet")

	// load mesh
	m, err := Read("data/cubeandtet.msh")
	if err != nil {
		tst.Errorf("Read failed:\n%v\n", err)
		return
	}

	// correct data
	nverts := 9
	ncells := 2
	allX := [][]float64{
		{0.0, 0.0, 0.0},
		{1.0, 0.0, 0.0},
		{1.0, 1.0, 0.0},
		{0.0, 1.0, 0.0},
		{0.0, 0.0, 1.0},
		{1.0, 0.0, 1.0},
		{1.0, 1.0, 1.0},
		{0.0, 1.0, 1.0},
		{0.0, 2.0, 0.0},
	}
	allVtags := []int{0, 0, -12, 0, -14, 0, 0, 0, -18}
	allCtags := []int{-1, -1}
	allParts := []int{0, 0}
	allTypeKeys := []string{"hex8", "tet4"}
	allTypeIndices := []int{TypeHex8, TypeTet4}
	allV := [][]int{
		{0, 1, 2, 3, 4, 5, 6, 7},
		{3, 2, 8, 7},
	}
	allEtags := [][]int{
		{-10, -11, -12, -13, 0, 0, 0, 0, 0, 0, -15, 0},
		{-12, -12, -12, 0, -66, 0},
	}

	// check input data
	checkinput(tst, m, nverts, ncells, allX, allVtags, allCtags, allParts, allTypeKeys, allTypeIndices, allV, allEtags, nil)

	// derived data
	ndim := 3
	xmin := []float64{0, 0, 0}
	xmax := []float64{1, 2, 1}
	allGndim := []int{3, 3}
	allCoords := [][][]float64{
		[][]float64{
			{0.0, 0.0, 0.0},
			{1.0, 0.0, 0.0},
			{1.0, 1.0, 0.0},
			{0.0, 1.0, 0.0},
			{0.0, 0.0, 1.0},
			{1.0, 0.0, 1.0},
			{1.0, 1.0, 1.0},
			{0.0, 1.0, 1.0},
		},
		[][]float64{
			{0.0, 1.0, 0.0},
			{1.0, 1.0, 0.0},
			{0.0, 2.0, 0.0},
			{0.0, 1.0, 1.0},
		},
	}

	// check derived data
	checkderived(tst, m, ndim, xmin, xmax, allGndim, allCoords)

	// get map of tags
	tm, err := m.GetTagMaps()
	if err != nil {
		tst.Errorf("GetTagMaps failed:\n%v\n", err)
		return
	}

	// correct data
	vtags := []int{-12, -14, -18}
	ctags := []int{-1}
	cparts := []int{0}
	etags := []int{-10, -11, -12, -13, -15, -66}
	ftags := []int{-100, -101, -200, -300, -400}
	ctypeinds := []int{TypeHex8, TypeTet4}
	vtagsVids := [][]int{{2}, {4}, {8}}
	ctagsCids := [][]int{{0, 1}}
	cpartsCids := [][]int{{0, 1}}
	ctypesCids := [][]int{{0}, {1}}
	etagsCids := [][]int{{0}, {0}, {0, 1, 1, 1}, {0}, {0}, {1}} // not unique
	etagsLocEids := [][]int{{0}, {1}, {2, 0, 1, 2}, {3}, {10}, {4}}
	etagsVids := [][]int{{0, 1}, {1, 2}, {2, 3, 8}, {0, 3}, {2, 6}, {2, 7}}
	ftagsCids := [][]int{{0, 1}, {0}, {0}, {0, 1}, {1}}
	ftagsLocEids := [][]int{{0, 0}, {1}, {2}, {4, 2}, {3}}
	ftagsVids := [][]int{{0, 3, 4, 7, 8}, {1, 2, 5, 6}, {0, 1, 4, 5}, {0, 1, 2, 3, 8}, {2, 7, 8}}

	// check maps
	checkmaps(tst, m, tm, vtags, ctags, cparts, etags, ftags, ctypeinds, vtagsVids, ctagsCids, cpartsCids, ctypesCids, etagsVids, etagsCids, etagsLocEids, ftagsVids, ftagsCids, ftagsLocEids)

	// draw
	if chk.Verbose {
		args := NewArgs()
		args.WithIdsCells = true
		args.WithEdges = true
		plt.Reset(true, nil)
		m.Draw(args)
		plt.HideAllBorders()
		plt.Save("/tmp/gosl/gm", "cubeandtet")
	}
}

func checkinput(tst *testing.T, m *Mesh, nverts, ncells int, X [][]float64, vtags, ctags, parts []int, typekeys []string, typeindices []int, V [][]int, etags, ftags [][]int) {
	if len(m.Verts) != nverts {
		tst.Errorf("nverts is incorrect: %d != %d", len(m.Verts), nverts)
		return
	}
	if len(m.Cells) != ncells {
		tst.Errorf("ncells is incorrect: %d != %d", len(m.Cells), ncells)
		return
	}
	io.Pfyel("\nvertices:\n")
	for i, v := range m.Verts {
		io.Pf("%+v\n", v)
		chk.Vector(tst, io.Sf("vertex %2d: X", v.Id), 1e-15, v.X, X[v.Id])
		if v.Tag != vtags[i] {
			tst.Errorf("vtag is incorrect: %d != %d", v.Tag, vtags[i])
			return
		}
	}
	io.Pfyel("\ncells:\n")
	for i, c := range m.Cells {
		io.Pf("%+v\n", c)
		if c.Tag != ctags[i] {
			tst.Errorf("ctag is incorrect: %d != %d", c.Tag, ctags[i])
			return
		}
		if c.Part != parts[i] {
			tst.Errorf("part is incorrect: %d != %d", c.Part, parts[i])
			return
		}
		chk.String(tst, typekeys[i], c.TypeKey)
		chk.Int(tst, "cell type index", typeindices[i], c.TypeIndex)
		chk.Ints(tst, io.Sf("cell %2d : V", c.Id), c.V, V[c.Id])
		chk.Ints(tst, io.Sf("cell %2d : edgetags", c.Id), c.EdgeTags, etags[c.Id])
	}
}

func checkderived(tst *testing.T, m *Mesh, ndim int, xmin, xmax []float64, allGndim []int, allCoords [][][]float64) {
	io.Pfyel("\nderived data:\n")
	chk.Int(tst, "Ndim", m.Ndim, ndim)
	chk.Vector(tst, "Xmin", 1e-15, m.Xmin, xmin)
	chk.Vector(tst, "Xmax", 1e-15, m.Xmax, xmax)
	for i, c := range m.Cells {
		chk.Matrix(tst, io.Sf("Cell %d: X", i), 1e-15, c.X, allCoords[i])
		chk.Int(tst, io.Sf("Cell %d: Gndim", i), c.Gndim, allGndim[i])
	}
}

func checkmaps(tst *testing.T, m *Mesh, tm *TagMaps, vtags, ctags, cparts, etags, ftags []int, ctypeinds []int, vtagsVids, ctagsCids, cpartsCids, ctypesCids, etagsVids, etagsCids, etagsLocEids, ftagsVids, ftagsCids, ftagsLocEids [][]int) {

	// VertTag2verts
	io.Pfyel("\nVertTag2verts:\n")
	for key, val := range tm.VertTag2verts {
		io.Pf("%v:\n", key)
		for _, s := range val {
			io.Pf("  vert: %v\n", s)
		}
	}
	if len(tm.VertTag2verts) != len(vtags) {
		tst.Errorf("size of map of vert tags is incorrect. %d != %d", len(tm.VertTag2verts), len(vtags))
		return
	}
	for i, tag := range vtags {
		var ids []int
		if verts, ok := tm.VertTag2verts[tag]; ok {
			for _, v := range verts {
				ids = append(ids, v.Id)
			}
		} else {
			tst.Errorf("cannot find tag %d in VertTag2verts map", tag)
			return
		}
		chk.Ints(tst, io.Sf("%d vertices", tag), ids, vtagsVids[i])
	}

	// CellTag2cells
	io.Pfyel("\nCellTag2cells:\n")
	for key, val := range tm.CellTag2cells {
		io.Pf("%v:\n", key)
		for _, s := range val {
			io.Pf("  cell: %v\n", s)
		}
	}
	if len(tm.CellTag2cells) != len(ctags) {
		tst.Errorf("size of map of cell tags is incorrect. %d != %d", len(tm.CellTag2cells), len(ctags))
		return
	}
	for i, tag := range ctags {
		var ids []int
		if cells, ok := tm.CellTag2cells[tag]; ok {
			for _, v := range cells {
				ids = append(ids, v.Id)
			}
		} else {
			tst.Errorf("cannot find tag %d in CellTag2cells map", tag)
			return
		}
		chk.Ints(tst, io.Sf("%d cells", tag), ids, ctagsCids[i])
	}

	// CellPart2cells
	io.Pfyel("\nCellPart2cells:\n")
	for key, val := range tm.CellPart2cells {
		io.Pf("%v:\n", key)
		for _, s := range val {
			io.Pf("  cell: %v\n", s)
		}
	}
	if len(tm.CellPart2cells) != len(cparts) {
		tst.Errorf("size of map of cell tags is incorrect. %d != %d", len(tm.CellPart2cells), len(cparts))
		return
	}
	for i, part := range cparts {
		var ids []int
		if cells, ok := tm.CellPart2cells[part]; ok {
			for _, v := range cells {
				ids = append(ids, v.Id)
			}
		} else {
			tst.Errorf("cannot find part %d in CellPart2cells map", part)
			return
		}
		chk.Ints(tst, io.Sf("%d cells", part), ids, cpartsCids[i])
	}

	// CellType2cells
	io.Pfyel("\nCellType2cells:\n")
	for key, val := range tm.CellType2cells {
		io.Pf("%v:\n", key)
		for _, s := range val {
			io.Pf("  cell: %v\n", s)
		}
	}
	if len(tm.CellType2cells) != len(ctypeinds) {
		tst.Errorf("size of map of cell tags is incorrect. %d != %d", len(tm.CellType2cells), len(ctypeinds))
		return
	}
	for i, typ := range ctypeinds {
		var ids []int
		if cells, ok := tm.CellType2cells[typ]; ok {
			for _, v := range cells {
				ids = append(ids, v.Id)
			}
		} else {
			tst.Errorf("cannot find type %q in CellType2cells map", TypeIndexToKey[typ])
			return
		}
		chk.Ints(tst, io.Sf("%q cells", typ), ids, ctypesCids[i])
	}

	// EdgeTag2cells
	io.Pfyel("\nEdgeTag2cells:\n")
	for key, val := range tm.EdgeTag2cells {
		io.Pf("%v:\n", key)
		for _, s := range val {
			io.Pf("  cell: %v\n", s)
		}
	}
	if len(tm.EdgeTag2cells) != len(etags) {
		tst.Errorf("size of map of edge tags (cells) is incorrect. %d != %d", len(tm.EdgeTag2cells), len(etags))
		return
	}
	for i, tag := range etags {
		var cids []int
		var bryids []int
		if pairs, ok := tm.EdgeTag2cells[tag]; ok {
			for _, pair := range pairs {
				cids = append(cids, pair.C.Id)
				bryids = append(bryids, pair.BryId)
			}
		} else {
			tst.Errorf("cannot find tag %d in EdgeTag2cells map", tag)
			return
		}
		chk.Ints(tst, io.Sf("%d edges => cells  ", tag), cids, etagsCids[i])
		chk.Ints(tst, io.Sf("%d edges => bry ids", tag), bryids, etagsLocEids[i])
	}

	// EdgeTag2verts
	io.Pfyel("\nEdgeTag2verts:\n")
	for key, val := range tm.EdgeTag2verts {
		io.Pf("%v:\n", key)
		for _, s := range val {
			io.Pf("  vert: %v\n", s)
		}
	}
	if len(tm.EdgeTag2verts) != len(etags) {
		tst.Errorf("size of map of edge tags (verts) is incorrect. %d != %d", len(tm.EdgeTag2verts), len(etags))
		return
	}
	for i, tag := range etags {
		var ids []int
		if verts, ok := tm.EdgeTag2verts[tag]; ok {
			for _, v := range verts {
				ids = append(ids, v.Id)
			}
		} else {
			tst.Errorf("cannot find tag %d in EdgeTag2verts map", tag)
			return
		}
		chk.Ints(tst, io.Sf("%d edges => verts", tag), ids, etagsVids[i])
	}

	// FaceTag2cells
	io.Pfyel("\nFaceTag2cells:\n")
	for key, val := range tm.FaceTag2cells {
		io.Pf("%v:\n", key)
		for _, s := range val {
			io.Pf("  cell: %v\n", s)
		}
	}
	if len(tm.FaceTag2cells) != len(ftags) {
		tst.Errorf("size of map of face tags (cells) is incorrect. %d != %d", len(tm.FaceTag2cells), len(ftags))
		return
	}
	for i, tag := range ftags {
		var cids []int
		var bryids []int
		if pairs, ok := tm.FaceTag2cells[tag]; ok {
			for _, pair := range pairs {
				cids = append(cids, pair.C.Id)
				bryids = append(bryids, pair.BryId)
			}
		} else {
			tst.Errorf("cannot find tag %d in FaceTag2cells map", tag)
			return
		}
		chk.Ints(tst, io.Sf("%d faces => cells  ", tag), cids, ftagsCids[i])
		chk.Ints(tst, io.Sf("%d faces => bry ids", tag), bryids, ftagsLocEids[i])
	}

	// FaceTag2verts
	io.Pfyel("\nFaceTag2verts:\n")
	for key, val := range tm.FaceTag2verts {
		io.Pf("%v:\n", key)
		for _, s := range val {
			io.Pf("  vert: %v\n", s)
		}
	}
	if len(tm.FaceTag2verts) != len(ftags) {
		tst.Errorf("size of map of face tags (verts) is incorrect. %d != %d", len(tm.FaceTag2verts), len(ftags))
		return
	}
	for i, tag := range ftags {
		var ids []int
		if verts, ok := tm.FaceTag2verts[tag]; ok {
			for _, v := range verts {
				ids = append(ids, v.Id)
			}
		} else {
			tst.Errorf("cannot find tag %d in FaceTag2verts map", tag)
			return
		}
		chk.Ints(tst, io.Sf("%d faces => verts", tag), ids, ftagsVids[i])
	}
}
