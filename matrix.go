package matrix

import (
	"fmt"
	"rand"
)

type Vector []int
type Matrix struct {
	n int
	v []int
}

func NewMatrix(n int) (m *Matrix) {
	m = new(Matrix)
	m.n = n
	m.v = make([]int, n*n)
	return m
}

func (m Matrix) Get(i int, j int) int {
	return m.v[i+j*m.n]
}

func (m Matrix) Set(i int, j int, v int) {
	m.v[i+j*m.n] = v
}

func (p Vector) Swap(i int, j int) {
	x := p[i]
	p[i] = p[j]
	p[j] = x
}

func (v Vector) Copy(w Vector) {
	for i := 0; i < len(v); i++ {
		v[i] = w[i]
	}
}

func (v Vector) Print() {
	for i := 0; i < len(v); i++ {
		fmt.Printf("%d ", v[i])
	}
	fmt.Print("\n")
}

func (m *Matrix) Print() {
	for i := 0; i < m.n; i++ {
		for j := 0; j < m.n; j++ {
			fmt.Printf("%d ", m.Get(i, j))
		}
		fmt.Print("\n")
	}
}

func Perm(p Vector) {
	n := len(p)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	for i := 0; i < n; i++ {
		p.Swap(i, i+rand.Intn(n-i))
	}
}

