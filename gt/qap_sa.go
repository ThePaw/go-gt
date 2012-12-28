// Copyright 2012 The Gt Authors. All rights reserved. See the LICENSE file.

package gt

// Solves the Quadratic Assignment problem using Simulated Annealing. 
// D. T. Connoly, 1990. An improved annealing scheme for the QAP. European Journal of Operational Research 46: 93-100.
// Based on implementation of E. Taillard, 1998. 
// WARNING: All SA algorithms have bad performance for QAP. See Duman & Orb 2007, Table 3.

import (
	"fmt"
	"math"
	"math/rand"
)

func initQAP(a *Matrix, b *Matrix, w Vector, c int64) (int64, int64, int64) {
	var (
		dmin, dmax int64
	)
	n := w.Len()
	for i := 0; i < 10000; i++ {
		r := rand.Int63n(n)
		s := rand.Int63n(n - 1)
		if s >= r {
			s = s + 1
		}
		d := delta(a, b, w, r, s)
		c += d
		dmin = min(dmin, d)
		dmax = max(dmax, d)
		w.Swap(r, s)
	}
	return c, dmin, dmax
}

// QAP_SolveSA solves the Quadratic Assignment Problem using Simulated Annealing. 
func QAP_SolveSA(a *Matrix, b *Matrix, p Vector, m int64) int64 {
	var i int64
	n := p.Len()
	w := make(Vector, n)
	w.Copy(p)
	cc := cost(a, b, p)
	c, dmin, dmax := initQAP(a, b, w, cc)
	var t0 float64 = float64(dmin + (dmax-dmin)/10.0)
	tf := float64(dmin)
	beta := (t0 - tf) / (float64(m) * t0 * tf)
	var fail int64 = 0
	tries := n * (n - 1) / 2
	tfound := t0
	var temp float64 = t0
	var r int64 = 0
	var s int64 = 1

	// SA iterations
	for i = 0; i < m; i++ {
		temp /= (beta*temp + 1)
		s++
		if s >= n {
			r++
			if r >= n-1 {
				r = 0
			}
			s = r + 1
		}
		d := delta(a, b, w, r, s)
		if (d < 0) || (rand.Float64() < math.Exp(-float64(d)/temp)) || (fail == tries) {
			c += d
			w.Swap(r, s)
			fail = 0
		} else {
			fail++
		}
		if fail == tries {
			beta = 0
			temp = tfound
		}

		// Best solution improved ?
		if c < cc {
			cc = c
			p.Copy(w)
			tfound = temp
			if Verbose {
				fmt.Printf("iteration %d: cost=%d\n", i, cc)
				p.Print()
			}
		}
	}
	return cc
}
