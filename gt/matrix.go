package gt

import (
	"fmt"
	"math/rand"
	"bufio"
	"strconv"
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

func skip(rd *bufio.Reader) {
	var b byte = ' '
	var err error
	for b == ' ' || b == '\t' || b == '\n' {
		b, err = rd.ReadByte()
		if err != nil {
			return
		}
	}
	rd.UnreadByte()
}

func wskip(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\t' {
			return s[i:]
		}
	}
	return ""
}

func end(s string) (i int) {
	for i = 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n'{
			return i
		}
	}
	return 0
}

func readInt(s string) (x int, i int){
	i = end(s)
	x, _ = strconv.Atoi(s[:i])
	return x, i
}

func readMatrix(rd *bufio.Reader, n int) *Matrix {
	M := NewMatrix(n)
	for i := 0; i < n; i++ {
		skip(rd)
		line, _ := rd.ReadString('\n')
		for j := 0; j < n; j++ {
			line = wskip(line)
			x, p := readInt(line)
			M.Set(j, i, x)
			if p == 0 {
				panic("bad integer")
			}
			line = line[p:]
		}
	}
	return M
}
