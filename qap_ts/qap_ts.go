// Copyright 2012 The Gt Authors. All rights reserved. See the LICENSE file.

// Solves the Quadratic Assignment Problem using the Robust taboo search. 
// E. D. Taillard 1991. Robust taboo search for the quadratic assignment problem. Parallel Computing 17, 1991: 443-455.
package main

import (
	. "code.google.com/p/go-gt/gt"
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		tabu_duration, aspiration, cost int64
	)
	cost = 999999999
	k := flag.Int("k", 2, "Number of resolutions")
	m := flag.Int("m", 10000, "Number of iterations")
	o := flag.Int("o", 0, "Opt-parameter")
	verbose := flag.Bool("v", false, "Verbose")
	flag.Parse()
	iter := int64(*m)
	opt := int64(*o)
	in := os.Stdin
	if flag.NArg() > 0 {
		file := flag.Arg(0)
		var err error
		in, err = os.Open(file)
		if in == nil {
			fmt.Printf("can't open file %s: %s\n", file, err)
			return
		}
	}
	n, a, b := Load(in)
	in.Close()
	tabu_duration = 8 * n
	aspiration = n * n * 5

	p := make(Vector, n)
	best_sol := make(Vector, n)
	for i := 0; i < *k; i++ {
		Perm(p)
		cc := QAP_SolveTS(a, b, p, opt, tabu_duration, aspiration, iter, *verbose)
		if cc < cost {

			cost = cc
			best_sol.Copy(p)

		}
	}
	if !*verbose {
		best_sol.Print()
	}
}
