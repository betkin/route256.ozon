package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAutoCorrect(t *testing.T) {
	var caseList [5]string = [5]string{
		"hello",
		"hello, it's me. you're looking for!",
		"one? two! Zero",
		".hello",
		""}
	var expectList [5]string = [5]string{
		"Hello.",
		"Hello, it's me. You're looking for!",
		"One? Two! Zero.",
		".Hello.",
		""}

	for i := 0; i < len(caseList); i++ {
		assert.Equal(t, expectList[i], AutoCorrect(caseList[i]))
	}
}
