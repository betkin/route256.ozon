package task2

import (
	"sort"
	"testing"
)

func Median(digits []float32) float32 {
	sort.Slice(digits, func(i, j int) bool {
		return digits[i] < digits[j]
	})
	if len(digits)%2 == 0 {
		return (digits[len(digits)/2-1] + digits[len(digits)/2]) / 2
	} else {
		return digits[len(digits)/2]
	}
}

func Avg(digits []float32) float32 {
	var amount float64 = 0
	for _, d := range digits {
		amount += float64(d)
	}
	return float32(amount / float64(len(digits)))
}

func Task2(digits []float32) (float32, float32) {
	return Avg(digits), Median(digits)
}

func TestAvg(t *testing.T) {
	var r [2]float32
	var caseList [5][]float32 = [5][]float32{
		{10},
		{99, 66},
		{1, 2, 3, 4, 5},
		{3, -1, 0, -4, 3, 10},
		{-1, -4, -5, -7, -0.000000001}}
	var resultList [5][2]float32 = [5][2]float32{
		{10, 10},
		{82.5, 82.5},
		{3, 3},
		{1.833333333333333, 1.5},
		{-3.4000000002, -4}}

	for i := 0; i < 5; i++ {
		r[0], r[1] = Task2(caseList[i])
		if r[0] != resultList[i][0] || r[1] != resultList[i][1] {
			t.Error("Expected ", resultList[i][0], resultList[i][1])
		}
	}
}
