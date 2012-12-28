// Copyright 2012 The Gt Authors. All rights reserved. See the LICENSE file.

package gt

// Solve the Quadratic Assignment Problem using the Fast ant system. 
// E. D. Taillard 1998. "FANT: Fast ant system.  Technical report IDSIA-46-98, IDSIA, Lugano.

import (
	"fmt"
)

// Local search: Scan the neighbourhood at most twice. 
// Perform improvements as soon as they are found. 
func localSearch(a *Matrix, b *Matrix, p Vector, cost *int64) {
	// set of moves, numbered from 0 to index
	var i, j, nMov int64
	n := p.Len()
	move := make(Vector, n*(n-1)/2)
	nMov = 0
	for i = 0; i < n-1; i++ {
		for j = i + 1; j < n; j++ {
			move[nMov] = n*i + j
			nMov++
		}
	}
	improved := true
	for k := 0; k < 2 && improved; k++ {
		improved = false
		for i = 0; i < nMov-1; i++ {
			move.Swap(i, unif(i+1, nMov-1))
		}
		for i = 0; i < nMov; i++ {
			r := move[i] / n
			s := move[i] % n
			d := delta(a, b, p, r, s)
			if d < 0 {
				*cost += d
				p.Swap(r, s)
				improved = true
			}
		}
	}
	return
}

// (Re-) initialization of the trace. 
func initTrace(n, inc int64, trace *Matrix) {
	var i, j int64
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			trace.Set(i, j, inc)
		}
	}
}

// Trace update. 
func updateTrace(n int64, p, best_p Vector, inc *int64, r int64, trace *Matrix) {
	var i int64
	for i = 0; i < n && p[i] == best_p[i]; i++ {  // skip
	}
	if i == n {
		(*inc)++
		initTrace(n, *inc, trace)
	} else {
		for i = 0; i < n; i++ {
			trace.Set(i, p[i], trace.Get(i, p[i])+*inc)
			trace.Set(i, best_p[i], trace.Get(i, best_p[i])+r)
		}
	}
}

// Generate a solution with probability of setting p[i] == j 
// proportionnal to trace[i][j]. 
func genTrace(p Vector, trace *Matrix) {
	var i, j, k, target, sum int64
	n := p.Len()
	nexti := make(Vector, n)
	nextj := make(Vector, n)
	sum_trace := make(Vector, n)

	Perm(nexti)
	Perm(nextj)
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			sum_trace[i] += trace.Get(i, j)
		}
	}

	for i = 0; i < n; i++ {
		target = unif(0, sum_trace[nexti[i]]-1)
		j = i
		sum = trace.Get(nexti[i], nextj[j])
		for sum < target {
			j++
			sum += trace.Get(nexti[i], nextj[j])
		}
		p[nexti[i]] = nextj[j]
		for k = i; k < n; k++ {
			sum_trace[nexti[k]] -= trace.Get(nexti[k], nextj[j])
		}
		nextj.Swap(j, i)
	}
}

// QAP_SolveFANT solves the Quadratic Assignment Problem using Fast Ant System. 
func QAP_SolveFANT(a *Matrix, b *Matrix, p Vector, r, m int64) int64 {
	var inc, i, c int64
	n := p.Len()
	w := make(Vector, n)
	w.Copy(p)
	trace := NewMatrix(n)
	inc = 1
	initTrace(n, inc, trace)
	cc := Inf

	// FANT iterations
	for i = 0; i < m; i++ {
		// Build a new solution
		genTrace(w, trace)
		c = cost(a, b, w)
		// Improve solution with a local search
		localSearch(a, b, w, &c)
		// Best solution improved ?
		if c < cc {
			cc = c
			p.Copy(w)
			if Verbose {
				fmt.Printf("iteration %d: cost=%d\n", i, cc)
				p.Print()
			}
			inc = 1
			initTrace(n, inc, trace)
		} else {
			// Memory update
			updateTrace(n, w, p, &inc, r, trace)
		}
	}
	return cc
}
