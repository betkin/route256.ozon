package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func AutoCorrect(txt string) string {
	var (
		worker  string
		dotFlag bool = true
	)

	txt = strings.TrimSpace(txt)
	if len(txt) == 0 {
		return ""
	}
	for _, ch := range txt {
		if dotFlag && unicode.IsLetter(ch) {
			dotFlag = false
			ch = unicode.ToUpper(ch)
		}
		if ch == 33 || ch == 46 || ch == 63 || ch == 10 {
			dotFlag = true
		}
		worker += string(ch)
	}
	worker = strings.TrimSpace(worker)
	var ch byte = txt[len(txt)-1]
	if ch != 33 && ch != 46 && ch != 63 || ch == 10 {
		worker += "."
	}
	return worker
}

func main() {
	var s string
	fmt.Println("Please, enter sentences for test.")
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Println("\nResult:")
	fmt.Print(AutoCorrect(s))
}
