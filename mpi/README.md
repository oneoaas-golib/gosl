# Gosl. mpi. Message Passing Interface for parallel computing

[![GoDoc](https://godoc.org/github.com/cpmech/gosl/mpi?status.svg)](https://godoc.org/github.com/cpmech/gosl/mpi) 

More information is available in **[the documentation of this package](https://godoc.org/github.com/cpmech/gosl/mpi).**

The `mpi` package is a light wrapper to the [OpenMPI](https://www.open-mpi.org) C++ library designed
to develop algorithms for parallel computing.

This package allows parallel computations over the network and extends the concurrency capabilities of Go.

Both `goroutines` and MPI calls can co-exist to assist on High Performance Computing (HPC) work.

## Examples

### Communication between 3 processors

The next code can be executed with the following command:
```bash
mpirun -np 3 go run my_mpi_code.go
```

```go
func setslice(x []float64) {
	switch mpi.Rank() {
	case 0:
		copy(x, []float64{0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3})
	case 1:
		copy(x, []float64{10, 10, 10, 20, 20, 20, 30, 30, 30, 40, 40})
	case 2:
		copy(x, []float64{100, 100, 100, 1000, 1000, 1000, 2000, 2000, 2000, 3000, 3000})
	}
}

mpi.Start(false)
defer mpi.Stop(false)

if mpi.Rank() == 0 {
    io.PfYel("\nTest MPI 01\n")
}
if mpi.Size() != 3 {
    chk.Panic("this test needs 3 processors")
}
n := 11
x := make([]float64, n)
id, sz := mpi.Rank(), mpi.Size()
start, endp1 := (id*n)/sz, ((id+1)*n)/sz
for i := start; i < endp1; i++ {
    x[i] = float64(i)
}

// Barrier
mpi.Barrier()

io.Pfgrey("x @ proc # %d = %v\n", id, x)

// SumToRoot
r := make([]float64, n)
mpi.SumToRoot(r, x)
var tst testing.T
if id == 0 {
    chk.Vector(&tst, fmt.Sprintf("SumToRoot:       r @ proc # %d", id), 1e-17, r, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
} else {
    chk.Vector(&tst, fmt.Sprintf("SumToRoot:       r @ proc # %d", id), 1e-17, r, make([]float64, n))
}

// BcastFromRoot
r[0] = 666
mpi.BcastFromRoot(r)
chk.Vector(&tst, fmt.Sprintf("BcastFromRoot:   r @ proc # %d", id), 1e-17, r, []float64{666, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

// AllReduceSum
setslice(x)
w := make([]float64, n)
mpi.AllReduceSum(x, w)
chk.Vector(&tst, fmt.Sprintf("AllReduceSum:    w @ proc # %d", id), 1e-17, w, []float64{110, 110, 110, 1021, 1021, 1021, 2032, 2032, 2032, 3043, 3043})

// AllReduceSumAdd
setslice(x)
y := []float64{-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000}
mpi.AllReduceSumAdd(y, x, w)
chk.Vector(&tst, fmt.Sprintf("AllReduceSumAdd: y @ proc # %d", id), 1e-17, y, []float64{-890, -890, -890, 21, 21, 21, 1032, 1032, 1032, 2043, 2043})

// AllReduceMin
setslice(x)
mpi.AllReduceMin(x, w)
chk.Vector(&tst, fmt.Sprintf("AllReduceMin:    x @ proc # %d", id), 1e-17, x, []float64{0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3})

// AllReduceMax
setslice(x)
mpi.AllReduceMax(x, w)
chk.Vector(&tst, fmt.Sprintf("AllReduceMax:    x @ proc # %d", id), 1e-17, x, []float64{100, 100, 100, 1000, 1000, 1000, 2000, 2000, 2000, 3000, 3000})
```
