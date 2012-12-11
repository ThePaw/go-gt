// Copyright 2012 The Gt Authors. All rights reserved. See the LICENSE file.

package gt

// Solve the Quadratic Assignment Problem. 

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
)

var verbose bool

func Load(in *os.File) (int64, *Matrix, *Matrix) {
	rd := bufio.NewReader(in)
	skip(rd)
	line, _ := rd.ReadString('\n')
	line = wskip(line)
	n, i := readUint(line)
	if i == 0 {
		panic("expecting Matrix size")
	}
	a := readMatrix(rd, n)
	b := readMatrix(rd, n)
	return n, a, b
}

func cost(a *Matrix, b *Matrix, p Vector) (c int64) {
	var i, j int64
	c = 0
	for i = 0; i < p.Len(); i++ {
		for j = 0; j < p.Len(); j++ {
			c += a.Get(i, j) * b.Get(p[i], p[j])
		}
	}
	return c
}

func delta(a *Matrix, b *Matrix, p Vector, r int64, s int64) (d int64) {
	var i int64
	d = int64((a.Get(r, r) - a.Get(s, s)) * (b.Get(p[s], p[s]) - b.Get(p[r], p[r]) + (a.Get(r, s)-a.Get(s, r))*(b.Get(p[s], p[r])-b.Get(p[r], p[s]))))
	for i = 0; i < p.Len(); i++ {
		if i != r && i != s {
			d += (a.Get(i, r)-a.Get(i, s))*(b.Get(p[i], p[s])-b.Get(p[i], p[r])) + (a.Get(r, i)-a.Get(s, i))*(b.Get(p[s], p[i])-b.Get(p[r], p[i]))
		}
	}
	return d
}

func initQAP(a *Matrix, b *Matrix, w Vector, c int64) (int64, int64, int64) {
	var (
		dmin, dmax int64
	)
	n := w.Len()
	for i := 0; i < 1000; i++ {
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

// Solve the Quadratic Assignment Problem using Simulated Annealing. 
func QAP_SolveSA(a *Matrix, b *Matrix, p, best_p Vector, m int64, verbose bool) int64 {
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
	k := n * (n - 1) / 2
	tries := k
	tfound := t0
	var temp float64 = t0
	var r int64 = 0
	var s int64 = 1
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
		if c < cc {
			cc = c
			p.Copy(w)
			tfound = temp
			if verbose {
				fmt.Printf("iteration %d: cost=%d\n", i, cc)
				p.Print()
			}
		}
	}
	return cc
}
