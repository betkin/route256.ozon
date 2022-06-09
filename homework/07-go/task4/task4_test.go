package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Sum row(x)

func Sumrow(sl [][]*int, x int) int {
	s := 0
	for i := 0; i < len(sl[x]); i++ {
		s += *sl[x][i]
	}
	return s
}

// Sum col(x)

func Sumcol(sl [][]*int, x int) int {
	s := 0
	for i := 0; i < len(sl); i++ {
		s += *sl[i][x]
	}
	return s
}

// Sums diagonal

func Sumsdiag(sl [][]*int) (int, int) {
	s1 := 0
	s2 := 0
	for i := 0; i < len(sl); i++ {
		s1 += *sl[i][i]
		s2 += *sl[len(sl)-i-1][i]
	}
	return s1, s2
}

func TestMagicsgen(t *testing.T) {
	const N = 3
	var m int = N * (N*N + 1) / 2
	var testgen = Magicsgen(N)

	for i := 0; i < N; i++ {
		assert.Equal(t, m, Sumrow(testgen, i), "Summa of rows are wrong")
		assert.Equal(t, m, Sumcol(testgen, i), "Summa of columns are wrong")
	}
	d1, d2 := Sumsdiag(testgen)
	assert.Equal(t, m, d1, "Summa of diagonal is wrong")
	assert.Equal(t, m, d2, "Summa of diagonal is wrong")
}
