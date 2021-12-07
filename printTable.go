package main

import (
	"math/rand"
	"time"
)

type getIndexer interface {
	getIndexes(int) []int
}

type showAllConfig struct {
}

func (sa showAllConfig) getIndexes(nRows int) []int {
	var n int
	if nRows > 0 {
		n = nRows
	}
	idxs := make([]int, n)
	for i := 0; i < n; i++ {
		idxs[i] = i
	}
	return idxs
}

type showHeadConfig struct {
	// Show up to this many rows from the start
	// of the table
	n int
}

func (sa showHeadConfig) getIndexes(nRows int) []int {
	var idxs []int
	for i := 0; i < nRows && i < sa.n; i++ {
		idxs = append(idxs, i)
	}
	return idxs
}

type showTailConfig struct {
	n int
}

func (sa showTailConfig) getIndexes(nRows int) []int {
	var idxs []int
	for i := 0; i < nRows && i < sa.n; i++ {
		idxs = append(idxs, nRows-i-1)
	}
	rev := make([]int, len(idxs))
	for i, j := range idxs {
		rev[len(rev)-i-1] = j
	}
	return rev
}

type showRandomConfig struct {
	n    int // Number of rows to select randomly
	seed int // RNG seed (Default: 0 will use current time)
}

func (sa showRandomConfig) getIndexes(nRows int) []int {
	// Create rng
	src := rand.NewSource(time.Now().UnixNano())
	if sa.seed != 0 {
		src.Seed(int64(sa.seed))
	}
	rng := rand.New(src)

	// Generate the numbers
	var idxs []int
	for i := 0; i < sa.n; i++ {
		n := rng.Intn(nRows)
		idxs = append(idxs, n)
	}
	return idxs
}
