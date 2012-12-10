// Copyright 2012 The Gt Authors. All rights reserved. See the LICENSE file.

package gt

// Solve the Quadratic Assignment Problem using the Fast ant system. 
// E. D. Taillard 1998. "FANT: Fast ant system.  Technical report IDSIA-46-98, IDSIA, Lugano.

import (
	"fmt"
	"math"
	"math/rand"
)

const infinite  int64 = math.MaxInt64

// Uniform random number. 
func unif(low, high int64) int64 {
	return low + int64(float64(high-low+1)*rand.Float64())
}

// Local search: Scan the neighbourhood at most twice. 
// Perform improvements as soon as they are found. 
func local_search(a *Matrix, b *Matrix, p Vector, cost *int64) {
	// set of moves, numbered from 0 to index
	var i, j, nr_moves int64
	n := p.Len()
	move := make(Vector, n*(n-1)/2)
	nr_moves = 0
	for i = 0; i < n-1; i++ {
		for j = i + 1; j < n; j++ {
			move[nr_moves] = n*i + j
			nr_moves++
		}
	}
	improved := true
	for scan_nr := 0; scan_nr < 2 && improved; scan_nr++ {
		improved = false
		for i = 0; i < nr_moves-1; i++ {
			move.Swap(i, unif(i+1, nr_moves-1))
		}
		for i = 0; i < nr_moves; i++ {
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

// (Re-) initialization of the memory. 
func init_trace(n, increment int64, trace *Matrix) {
	var i, j int64
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			trace.Set(i, j, increment)
		}
	}
}

// Memory update. 
func update_trace(n int64, p, best_p Vector, increment *int64, r int64, trace *Matrix) {
	var i int64
	for i = 0; i < n && p[i] == best_p[i]; i++ {
	}
	if i == n {
		(*increment)++
		init_trace(n, *increment, trace)
	} else {
		for i = 0; i < n; i++ {
			trace.Set(i, p[i], trace.Get(i, p[i])+*increment)
			trace.Set(i, best_p[i], trace.Get(i, best_p[i])+r)
		}
	}
}

// Generate a solution with probability of setting p[i] == j 
// proportionnal to trace[i][j]. 
func generate_solution_trace(n int64, p Vector, trace *Matrix) {
	var i, j, k, target, sum int64
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

// Solve the Quadratic Assignment Problem using Fast Ant System. 
func QAP_SolveFANT(a *Matrix, b *Matrix, p , best_p Vector, rr, mm int) int64 {
	var increment, i, k, r, m, cc int64
	m = int64(mm)
	r = int64(rr)
	n := p.Len()
	trace := NewMatrix(n)
	increment = 1
	init_trace(n, increment, trace)
	best_cost := infinite

	// FANT iterations
	for i = 0; i < m; i++ {
		// Build a new solution
		generate_solution_trace(n, p, trace)
		cc = cost(a, b, p)
		// Improve solution with a local search
		local_search(a, b, p, &cc)
		// Best solution improved ?
		if cc < best_cost {
			best_cost = cc
			if Verbose {
				fmt.Printf("iteration %d: cost=%d\n", i, cc)
				p.Print()
			}
			for k = 0; k < n; k = k + 1 {
				best_p[k] = p[k]
			}
			//      print(n, p);
			increment = 1
			init_trace(n, increment, trace)
		} else {
			// Memory update
			update_trace(n, p, best_p, &increment, r, trace)
		}
	}
	return cc
}
