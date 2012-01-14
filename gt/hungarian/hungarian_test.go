package hungarian

import (
	"gt"
	"testing"
	"fmt"
)

func assert(t bool) {
	if !t {
		panic("fail")
	}
}

func TestHungarian(t *testing.T) {
	g := new(gt.Matrix)
	g.N = 4
	g.A = []int64{
		9, 1, 9, 9,
		1, 9, 1, 9,
		9, 1, 9, 1,
		9, 9, 1, 9}
	p := Hungarian(g)
	fmt.Print(p)
	assert(p[0] == 0)
}
