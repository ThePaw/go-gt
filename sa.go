package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"strconv"
	"rand"
	"math"
)

type vector []int
type matrix struct {
	n int
	v []int
}

func newMatrix(n int) (m *matrix) {
	m = new(matrix)
	m.n = n
	m.v = make([]int, n*n)
	return m
}

var verbose bool

func (m matrix) get(i int, j int) int {
	return m.v[i+j*m.n]
}

func (m matrix) set(i int, j int, v int) {
	m.v[i+j*m.n] = v
}

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
	var err os.Error
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

func end(s string) int {
	for i := 0; i < len(s); i++ {
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

func readMatrix(rd *bufio.Reader, n int) *matrix {
	M := newMatrix(n)
	for i := 0; i < n; i++ {
		skip(rd)
		line, _ := rd.ReadString('\n')
		for j := 0; j < n; j++ {
			line = wskip(line)
			x, p := readInt(line)
			M.set(j, i, x)
			if p == 0 {
				panic("bad int64eger")
			}
			line = line[p:]
		}
	}
	return M
}

func load(in *os.File) (int, *matrix, *matrix) {
	rd := bufio.NewReader(in)
	skip(rd)
	line, _ := rd.ReadString('\n')
	line = wskip(line)
	n, i := readInt(line)
	if i == 0 {
		panic("expecting matrix size")
	}
	a := readMatrix(rd, n)
	b := readMatrix(rd, n)
	return n, a, b
}

func (p vector) swap(i int, j int) {
	x := p[i]
	p[i] = p[j]
	p[j] = x
}

func (v vector) copy(w vector) {
	for i := 0; i < len(v); i++ {
		v[i] = w[i]
	}
}

func (v vector) print() {
	for i := 0; i < len(v); i++ {
		fmt.Printf("%d ", v[i])
	}
	fmt.Print("\n")
}

func (m *matrix) print() {
	for i := 0; i < m.n; i++ {
		for j := 0; j < m.n; j++ {
			fmt.Printf("%d ", m.get(i, j))
		}
		fmt.Print("\n")
	}
}

func perm(p vector) {
	n := len(p)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	for i := 0; i < n; i++ {
		p.swap(i, i+rand.Intn(n-i))
	}
}

func cost(a *matrix, b *matrix, p vector) (c int64) {
	c = 0
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p); j++ {
			c += int64(a.get(i, j)*b.get(p[i], p[j]))
		}
	}
	return c
}

func delta(a *matrix, b *matrix, p vector, r int, s int) (d int64) {
	d = int64((a.get(r,r)-a.get(s,s))*(b.get(p[s],p[s])-b.get(p[r],p[r])) +
		(a.get(r,s)-a.get(s,r))*(b.get(p[s],p[r])-b.get(p[r],p[s])))
	for i := 0; i < len(p); i++ {
		if i != r && i != s {
			d += int64((a.get(i,r)-a.get(i,s))*(b.get(p[i],p[s])-b.get(p[i],p[r])) +
				(a.get(r,i)-a.get(s,i))*(b.get(p[s],p[i])-b.get(p[r],p[i])))
		}
	}
	return d

}

func inits(a *matrix, b *matrix, w vector, c int64) (int64, int64, int64) {
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
		w.swap(r, s)
	}
	return c, dmin, dmax
}

func solve(a *matrix, b *matrix, v vector, m int) int64 {
	n := len(v)
	w := make(vector, n)
	w.copy(v)
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
			w.swap(r, s)
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
			v.copy(w)
			tfound = temp
			if verbose {
				fmt.Printf("iteration %d: cost=%d\n", i, cc)
			}
		}
	}
	return cc
}

func main() {
	k := flag.Int("k", 1, "Number of resolutions")
	m := flag.Int("m", 1000, "Number of iterations")
	ver := flag.Bool("v", false, "Verbose")
	flag.Parse()
	verbose = *ver
	in := os.Stdin
	if len(os.Args) > 1 {
		file := flag.Arg(0)
		var err os.Error
		in, err = os.Open(file)
		if in == nil {
			fmt.Printf("can't open file %s: %s\n", file, err.String())
			return
		}
	}
	n, a, b := load(in)
	in.Close()
	v := make(vector, n)
	perm(v)
	for i := 0; i < *k; i++ {
		solve(a, b, v, *m)
	}
	v.print()
}
