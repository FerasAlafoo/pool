package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	s := os.Args[1]
	if s == "" {
		return
	}
	toPrint := ""
	end := len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			toPrint += s[i:end]
			end = i
		}
	}
	toPrint += " " + s[:end]
	for _, r := range toPrint {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
