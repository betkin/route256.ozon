package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromMaxToMin(t *testing.T) {
	var caseList = [][]float32{
		{13},
		{-5, -10},
		{10, 20, 30, 40, 50},
		{3, -1, 0, -4, 3, 10},
		{-1.5, -4.6, -5.1, 7, -0.000000001}}
	var expectList = []float32{
		0,
		5,
		40,
		14,
		12.1}

	for i := 0; i < len(caseList); i++ {
		assert.Equal(t, expectList[i], FromMaxToMin(caseList[i]))
	}
}
