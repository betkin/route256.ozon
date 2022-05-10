package main

import "testing"

func TestAutoCorrect(t *testing.T) {
	var r string
	var caseList [5]string = [5]string{
		"hello",
		"hello, it's me. you're looking for!",
		"one? two! Zero",
		".hello",
		""}
	var resultList [5]string = [5]string{
		"Hello.",
		"Hello, it's me. You're looking for!",
		"One? Two! Zero.",
		".Hello.",
		""}

	for i := 0; i < 5; i++ {
		r = AutoCorrect(caseList[i])
		if r != resultList[i] {
			t.Error("Expected ", resultList[i])
		}
	}
}
