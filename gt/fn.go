// Copyright 2012 The Gt Authors. All rights reserved. See the LICENSE file.

package gt

// Some handy functions. 

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
