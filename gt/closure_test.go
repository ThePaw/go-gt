package gt

import (
	"testing"
)

func assert(t bool) {
	if !t {
		panic("fail")
	}
}

func TestClosure(t *testing.T) {
	g := new(Matrix)
	g.N = 4
	g.A = []int64{
		0, 1, 1, 0,
		1, 0, 0, 0,
		1, 0, 0, 1,
		0, 0, 1, 0}
	c := NewMatrix(4)
	g.TransitiveClosure(c)
	assert(g.Get(0, 3) == 2)
}

func TestPath(t *testing.T) {
	g := new(Matrix)
	g.N = 4
	g.A = []int64{
		0, 1, 0, 0,
		1, 0, 1, 0,
		0, 1, 0, 1,
		0, 0, 1, 0}
	c := NewMatrix(4)
	g.TransitiveClosure(c)
	p := g.ShortestPath(0, 3, c)
	assert(p.Len() == 3)
}
