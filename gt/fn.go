// Copyright 2012 The Gt Authors. All rights reserved. See the LICENSE file.

package gt

// Some handy functions. 

import (
	"math"
	"math/rand"
)

var Verbose bool


const Inf int64 = math.MaxInt64

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func cube(x float64) float64 {
	return x * x * x
}

// Uniform random number. 
func unif(low, high int64) int64 {
	return low + int64(float64(high-low+1)*rand.Float64())
}
