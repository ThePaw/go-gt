package gt

import (
	"bufio"
	"fmt"
	"math/rand"
	"strconv"
)

type Vector []int64

type Matrix struct {
	N int64
	A []int64
}

func NewMatrix(n int64) (m *Matrix) {
	m = new(Matrix)
	m.N = n
	m.A = make([]int64, n*n)
	return m
}

func (m Matrix) Get(i int64, j int64) int64 {
	return m.A[i*m.N+j]
}

func (m Matrix) Set(i int64, j int64, v int64) {
	m.A[i*m.N+j] = v
}

func (p Vector) Swap(i int64, j int64) {
	x := p[i]
	p[i] = p[j]
	p[j] = x
}

func (v Vector) Len() int64 {
	return int64(len(v))
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
	var i, j int64
	for i = 0; i < m.N; i++ {
		for j = 0; j < m.N; j++ {
			fmt.Printf("%d ", m.Get(i, j))
		}
		fmt.Print("\n")
	}
}

func Perm(p Vector) {
	n := int64(len(p))
	var i int64
	for i = 0; i < n; i++ {
		p[i] = int64(i)
	}
	for i = 0; i < n; i++ {
		p.Swap(i, i+rand.Int63n(n-i))
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

func end(s string) (i int64) {
	for i = 0; i < int64(len(s)); i++ {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			return i
		}
	}
	return 0
}

func readUint(s string) (int64, int64) {
	i := end(s)
	x, _ := strconv.ParseInt(s[:i], 10, 64)
	return int64(x), i
}

func readMatrix(rd *bufio.Reader, n int64) *Matrix {
	M := NewMatrix(n)
	var i, j int64
	for i = 0; i < n; i++ {
		skip(rd)
		line, _ := rd.ReadString('\n')
		for j = 0; j < n; j++ {
			line = wskip(line)
			x, p := readUint(line)
			M.Set(j, i, x)
			if p == 0 {
				panic("bad int")
			}
			line = line[p:]
		}
	}
	return M
}
