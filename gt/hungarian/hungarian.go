// Hungarian algorithm to solve the assigment problem

package hungarian

import (
	"go-gt.googlecode.com/hg/gt"
	"math"
)

type env struct {
	m, n   int64
	g      *gt.Matrix
	T, S   []bool
	slack  []int64
	slackx []int64
	prev   []int64
	xy, yx []int64
	lx, ly []int64
}

func newEnv(n int64) *env {
	e := new(env)
	e.m = 0
	e.n = n
	e.T = make([]bool, n)
	e.S = make([]bool, n)
	e.slack = make([]int64, n)
	e.slackx = make([]int64, n)
	e.prev = make([]int64, n)
	e.xy = make([]int64, n)
	e.yx = make([]int64, n)
	e.lx = make([]int64, n)
	e.ly = make([]int64, n)
	var i int64
	for i = 0; i < n; i++ {
		e.xy[i] = -1
		e.yx[i] = -1
	}
	return e
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func (e *env) update() {
	var i int64
	var d int64 = math.MaxInt64
	for i = 0; i < e.n; i++ {
		if !e.T[i] {
			d = min(d, e.slack[i])
		}
	}
	for i = 0; i < e.n; i++ {
		if e.S[i] {
			e.lx[i] -= d
		}
	}
	for i = 0; i < e.n; i++ {
		if e.T[i] {
			e.ly[i] += d
		}
	}
	for i = 0; i < e.n; i++ {
		if !e.T[i] {
			e.slack[i] -= d
		}
	}
}

func (e *env) add(i, p int64) {
	var j int64
	e.S[i] = true
	e.prev[i] = p
	for j = 0; j < e.n; j++ {
		if e.lx[i]+e.ly[i]-e.g.Get(i, j) < e.slack[i] {
			e.slack[i] = e.lx[i] + e.ly[i] - e.g.Get(i, j)
			e.slackx[i] = j
		}
	}
}

func (e *env) augment() {
	var i, j, wr, rd, r int64
	wr = 0
	rd = 0
	q := make([]int64, e.n)
	if e.m == e.n {
		return
	}
	for i = 0; i < e.n; i++ {
		if e.xy[i] == -1 {
			wr++
			q[wr] = i
			r = i
			e.prev[i] = -2
			e.S[i] = true
			break
		}
	}
	for i = 0; i < e.n; i++ {
		e.slack[i] = e.lx[r] + e.ly[i] - e.g.Get(r, i)
		e.slackx[i] = r
	}
	for {
		for rd < wr {
			rd++
			i = q[rd]
			for j = 0; j < e.n; j++ {
				if e.g.Get(i, j) == e.lx[i]+e.ly[j] && !e.T[j] {
					if e.yx[j] == -1 {
						break
					}
					e.T[j] = true
					wr++
					q[wr] = e.yx[j]
					e.add(e.yx[j], i)
				}
			}
			if j < e.n {
				break
			}
		}
		if j < e.n {
			break
		}
		e.update()
		wr = 0
		rd = 0
		for j = 0; j < e.n; j++ {
			if !e.T[j] && e.slack[j] == 0 {
				if e.yx[i] == -1 {
					i = e.slackx[j]
					break
				} else {
					e.T[j] = true
					if !e.S[e.yx[j]] {
						wr++
						q[wr] = e.yx[j]
						e.add(e.yx[j], e.slackx[j])
					}
				}
			}
		}
		if j < e.n {
			return
		}
	}
	if j < e.n {
		e.m++
		for i != -2 {
			k := e.xy[i]
			e.yx[j] = i
			e.xy[i] = j
			i = e.prev[i]
			j = k
		}
		e.augment()
	}
}

func inits(g *gt.Matrix) (e *env) {
	var i, j int64
	e = newEnv(g.N)
	e.g = g
	e.n = g.N
	for i = 0; i < e.n; i++ {
		for j = 0; j < e.n; j++ {
			e.lx[i] = max(e.lx[i], e.g.Get(i, j))
		}
	}
	return e
}

func Hungarian(g *gt.Matrix) (xy, yx []int64) {
	e := inits(g)
	e.augment()
	return e.xy, e.yx
}
