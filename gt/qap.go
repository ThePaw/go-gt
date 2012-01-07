// Solve the Quadratic Assignment Problem
package gt

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"math/rand"
	"math"
)

var Verbose bool

func max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a int64, b int64) int64 {
	if a > b {
		return b
	}
	return a
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

func Load(in *os.File) (int, *Matrix, *Matrix) {
	rd := bufio.NewReader(in)
	skip(rd)
	line, _ := rd.ReadString('\n')
	line = wskip(line)
	n, i := readInt(line)
	if i == 0 {
		panic("expecting Matrix size")
	}
	a := readMatrix(rd, n)
	b := readMatrix(rd, n)
	return n, a, b
}

func cost(a *Matrix, b *Matrix, p Vector) (c int64) {
	c = 0
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p); j++ {
			c += int64(a.Get(i, j)*b.Get(p[i], p[j]))
		}
	}
	return c
}

func delta(a *Matrix, b *Matrix, p Vector, r int, s int) (d int64) {
	d = int64((a.Get(r,r)-a.Get(s,s))*(b.Get(p[s],p[s])-b.Get(p[r],p[r])) +
		(a.Get(r,s)-a.Get(s,r))*(b.Get(p[s],p[r])-b.Get(p[r],p[s])))
	for i := 0; i < len(p); i++ {
		if i != r && i != s {
			d += int64((a.Get(i,r)-a.Get(i,s))*(b.Get(p[i],p[s])-b.Get(p[i],p[r])) +
				(a.Get(r,i)-a.Get(s,i))*(b.Get(p[s],p[i])-b.Get(p[r],p[i])))
		}
	}
	return d

}

func inits(a *Matrix, b *Matrix, w Vector, c int64) (int64, int64, int64) {
	var (
		dmin, dmax int64
	)
	n := len(w)
	for i := 0; i < 1000; i++ {
		r := rand.Intn(n)
		s := rand.Intn(n-1)
		if s >= r {
			s = s+1
		}
		d := delta(a, b, w, r, s)
		c += d
		dmin = min(dmin, d)
		dmax = max(dmax, d)
		w.Swap(r, s)
	}
	return c, dmin, dmax
}

// Solve the Quadratic Assignment Problem using simulated annealing
func QAP_SolveSA(a *Matrix, b *Matrix, v Vector, m int) int64 {
	n := len(v)
	w := make(Vector, n)
	w.Copy(v)
	cc := cost(a, b, v)
	c, dmin, dmax := inits(a, b, w, cc)
	var t0 float64 = float64(dmin+(dmax-dmin)/10.0)
	tf := float64(dmin)
	beta := (t0-tf)/(float64(m)*t0*tf)
	fail := 0
	k := n*(n-1)/2
	tries := k
	tfound := t0
	var temp float64 = t0
	r := 0
	s := 1
	for i := 0; i < m; i++ {
		temp /= (beta*temp+1)
		s++
		if s >= n {
			r++
			if r >= n-1 {
				r = 0
			}
			s = r+1
		}
		d := delta(a, b, w, r, s)
		if (d < 0) || (rand.Float64() < math.Exp(-float64(d)/temp)) ||
				(fail == tries) {
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
			v.Copy(w)
			tfound = temp
			if Verbose {
				fmt.Printf("iteration %d: cost=%d\n", i, cc)
			}
		}
	}
	return cc
}

