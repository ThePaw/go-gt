// Copyright 2012 The Gt Authors. All rights reserved. See the LICENSE file.

package gt

// Common functions for  the Quadratic Assignment Problem.

import (
	"bufio"
	"os"
)

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
	d = int64((a.Get(r, r)-a.Get(s, s))*(b.Get(p[s], p[s])-b.Get(p[r], p[r])) +
		(a.Get(r, s)-a.Get(s, r))*(b.Get(p[s], p[r])-b.Get(p[r], p[s])))
	for i = 0; i < p.Len(); i++ {
		if i != r && i != s {
			d += (a.Get(i, r)-a.Get(i, s))*(b.Get(p[i], p[s])-b.Get(p[i], p[r])) +
				(a.Get(r, i)-a.Get(s, i))*(b.Get(p[s], p[i])-b.Get(p[r], p[i]))
		}
	}
	return d
}

// Cost difference if elements i and j  are swapped in permutation (solution) p, 
// but the value of dist[i][j] is supposed to
// be known before the transposition of elements r and s. 
func delta_part(a, b, dist *Matrix, p Vector, i, j, r, s int64) int64 {
	return (dist.Get(i, j) + (a.Get(r, i)-a.Get(r, j)+a.Get(s, j)-a.Get(s, i))*
		(b.Get(p[s], p[i])-b.Get(p[s], p[j])+b.Get(p[r], p[j])-b.Get(p[r], p[i])) +
		(a.Get(i, r)-a.Get(j, r)+a.Get(j, s)-a.Get(i, s))*
			(b.Get(p[i], p[s])-b.Get(p[j], p[s])+b.Get(p[j], p[r])-b.Get(p[i], p[r])))
}
