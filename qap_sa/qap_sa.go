// Copyright 2012 The Gt Authors. All rights reserved. See the LICENSE file.

// Solves the Quadratic Assignment problem using Simulated Annealing. 
package main

import (
	"flag"
	"fmt"
	"os"
	. "code.google.com/p/go-gt/gt"
)

var Verbose bool

func main() {
	k := flag.Int("k", 1, "Number of resolutions")
	m := flag.Int("m", 1000, "Number of iterations")
	verbose := flag.Bool("v", false, "Verbose")
	flag.Parse()
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
	best_p := make(Vector, n)
	Perm(p)
	best_p.Copy(p)
	for i := 0; i < *k; i++ {
		Perm(p)
		QAP_SolveSA(a, b, p, best_p, iter, *verbose)
	}
	if ! *verbose {
		best_p.Print()
	}
}
