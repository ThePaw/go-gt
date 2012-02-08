// Dijkstra's Algorithm for the shortest path.

package dijkstra

import (
	"code.google.com/p/go-gt/gt"
	"math"
)

type heap struct {
	n int64
	i []int64
	a []int64
	w []int64
}

func new(n int64) (h heap) {
	h.n = n
	h.i = make([]int64, n)
	h.a = make([]int64, n)
	h.w = make([]int64, n)
	var i int64
	for i = 0; i < n; i++ {
		h.i[i] = i
		h.a[i] = i
		h.w[i] = math.MaxInt64
	}
	return h
}

func (h heap) less(a, b int64) bool {
	i, j := h.i[a], h.i[b]
	return h.w[i] < h.w[j]
}

func (h heap) swap(a, b int64) {
	i, j := h.i[a], h.i[b]
	h.i[a], h.i[b] = h.i[b], h.i[a]
	h.a[i], h.a[j] = b, a
}

func (h heap) up(j int64) {
	for {
		i := (j - 1) / 2
		if i == j || h.less(i, j) {
			break
		}
		h.swap(i, j)
		j = i
	}
}

func (h heap) down(i int64) {
	for {
		left := 2*i + 1
		if left >= h.n {
			break
		}
		j := left
		if right := left + 1; right < h.n && !h.less(left, right) {
			j = right
		}
		if h.less(i, j) {
			break
		}
		h.swap(i, j)
		i = j
	}
}

func (h *heap) pop() (i int64) {
	i = h.i[0]
	h.n--
	h.swap(0, h.n)
	h.down(0)
	return i
}

func (h heap) update(p []int64, i int64, G *gt.Matrix) {
	var j int64
	for j = 0; j < G.N; j++ {
		if G.Get(i, j) > 0 {
			if h.w[i]+G.Get(i, j) < h.w[j] {
				p[j] = i + 1
				h.w[j] = h.w[i] + G.Get(i, j)
				h.up(h.a[j])
			}
		}
	}
}

func Dijkstra(G *gt.Matrix, i int64) (p []int64) {
	p = make([]int64, G.N)
	h := new(G.N)
	h.w[i] = 0
	h.swap(i, 0)
	for h.n > 0 {
		i = h.pop()
		if h.w[i] == math.MaxInt64 {
			return p
		}
		h.update(p, i, G)
	}
	return p
}
