package gt

import (
	"testing"
)

func assert(t bool) {
	if !t {
		panic("fail")
	}
}

func TestDijkstra(t *testing.T) {
	var g Matrix
	g.N = 4
	g.A = []int64{
		0, 1, 0, 0,
		1, 0, 1, 0,
		0, 1, 0, 1,
		0, 0, 1, 0}
	p := Dijkstra(&g, 0)
	assert(p[0] == 0)
	assert(p[1] == 1)
	assert(p[2] == 2)
	assert(p[3] == 3)
}
