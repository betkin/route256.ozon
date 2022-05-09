package task1

import (
	"testing"
	"unicode"
)

func AutoCorrect(txt string) string {
	var (
		l       int = len(txt)
		worker  string
		dotFlag bool = true
	)

	if l == 0 {
		return ""
	}
	for _, ch := range txt {
		if dotFlag && unicode.IsLetter(ch) {
			dotFlag = false
			ch = unicode.ToUpper(ch)
		}
		if ch == 33 || ch == 46 || ch == 63 {
			dotFlag = true
		}
		worker += string(ch)
	}
	var ch byte = txt[l-1]
	if ch != 33 && ch != 46 && ch != 63 {
		worker += "."
	}
	return worker
}

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
