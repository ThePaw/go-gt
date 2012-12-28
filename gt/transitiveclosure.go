// Copyright 2012 The Gt Authors. All rights reserved. See the LICENSE file.

package gt

import (
	"container/list"
)

// TransitiveClosure of a given graph G connects vertices u and v iff there is a path in G from u to v. Thus the transitive closure of any connected graph is complete.
func (G *Matrix) TransitiveClosure(N *Matrix) {
	var i, j, k int64
	for i = 0; i < G.N; i++ {
		for j = 0; j < G.N; j++ {
			for k = 0; k < G.N; k++ {
				if G.Get(i, k) > 0 && G.Get(k, j) > 0 {
					if G.Get(i, j) == 0 || G.Get(i, k)+G.Get(k, j) < G.Get(i, j) {
						G.Set(i, j, G.Get(i, k)+G.Get(k, j))
						if N != nil {
							N.Set(i, j, k+1)
						}
					}
				}
			}
		}
	}
}

// ShortestPath implements the algorithm for the shortest path.
func (G *Matrix) ShortestPath(src, tar int64, N *Matrix) (p *list.List) {
	p = list.New()
	if G.Get(src, tar) == 0 {
		return
	}
	next := N.Get(src, tar)
	if next == 0 {
		p.PushBack(tar)
	} else {
		p.PushBackList(G.ShortestPath(src, next-1, N))
		p.PushBackList(G.ShortestPath(next-1, tar, N))
	}
	return
}
