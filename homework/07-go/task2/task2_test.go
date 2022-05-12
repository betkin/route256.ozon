package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAvg(t *testing.T) {
	var r1, r2 float32
	var caseList = [][]float32{
		{10},
		{99, 66},
		{1, 2, 3, 4, 5},
		{3, -1, 0, -4, 3, 10},
		{-1, -4, -5, -7, -0.000000001}}
	var expectList = [][]float32{
		{10, 10},
		{82.5, 82.5},
		{3, 3},
		{1.833333333333333, 1.5},
		{-3.4000000002, -4}}

	for i := 0; i < len(caseList); i++ {
		r1, r2 = Task2(caseList[i])
		assert.Equal(t, expectList[i][0], r1, "Average value is wrong!")
		assert.Equal(t, expectList[i][1], r2, "Median is wrong!")
	}
}
