package main

import "testing"

func TestFromMaxToMin(t *testing.T) {
	var r float32
	var caseList [5][]float32 = [5][]float32{
		{13},
		{-5, -10},
		{10, 20, 30, 40, 50},
		{3, -1, 0, -4, 3, 10},
		{-1.5, -4.6, -5.1, 7, -0.000000001}}
	var resultList [5]float32 = [5]float32{
		0,
		5,
		40,
		14,
		12.1}

	for i := 0; i < 5; i++ {
		r = FromMaxToMin(caseList[i])
		if r != resultList[i] {
			t.Error("Expected ", resultList[i])
		}
	}
}
