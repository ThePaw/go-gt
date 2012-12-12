// Copyright 2012 The Gt Authors. All rights reserved. See the LICENSE file.

// Solves the Quadratic Assignment problem using the Fast Ant System. 
package main

import (
	. "code.google.com/p/go-gt/gt"
	"flag"
	"fmt"
	"os"
)

func main() {
	var cost = Inf
	k := flag.Int("k", 3, "Number of resolutions")
	m := flag.Int("m", 1000, "Number of iterations")
	rr := flag.Int("r", 5, "R-parameter")
	verbose := flag.Bool("v", false, "Verbose")
	flag.Parse()
	r := int64(*rr)
	iter := int64(*m)
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

	p := make(Vector, n)
	best_sol := make(Vector, n)
	for i := 0; i < *k; i++ {
		Perm(p)
		cc := QAP_SolveFANT(a, b, p, r, iter, *verbose)
		if cc < cost {
			cost = cc
			best_sol.Copy(p)
		}
	}
	if *verbose {
		fmt.Println("==============================")
		fmt.Println("best cost: ", cost)
	}
	best_sol.Print()
}
