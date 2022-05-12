package main

import "fmt"

func FromMaxToMin(digits []float32) float32 {
	var max, min float32
	max = digits[0]
	min = digits[0]
	for _, d := range digits {
		if max < d {
			max = d
		} else {
			if min > d {
				min = d
			}
		}
	}
	return max - min
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
	fmt.Print(FromMaxToMin(sl))
}
