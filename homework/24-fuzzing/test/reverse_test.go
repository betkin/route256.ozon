package test

import (
	"testing"

	fuzzing_example "gitlab.ozon.dev/betkin/device-api/homework/24-fuzzing/internal"
)

func MyIdealReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func FuzzReverse(f *testing.F) {
	f.Add("晓明晓明晓")
	f.Fuzz(func(t *testing.T, s string) {
		expect := fuzzing_example.MyReverse(s)
		answer := MyIdealReverse(s)
		if expect != answer {
			t.Errorf("%s IS NOT CORRECT RESULT FOR %s", expect, s)
		}
	})
}
