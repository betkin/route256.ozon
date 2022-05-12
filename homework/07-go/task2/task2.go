package main

import (
	"fmt"
	"sort"
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

func main() {
	var (
		err error
		x   float32
		sl  []float32
	)
	fmt.Println("Please, enter any numbers:")
	_, err = fmt.Scanf("%f", &x)
	for err == nil {
		sl = append(sl, x)
		_, err = fmt.Scanf("%f", &x)
	}
	fmt.Println("\nResult:")
	fmt.Print(Task2(sl))
}
