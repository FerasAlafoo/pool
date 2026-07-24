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

	ts := []string{}
	word := ""

	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			start = i
			break
		}
	}
	for i := start; i < len(s); i++ {
		if s[i] == ' ' && word != "" {
			ts = append(ts, word)
			word = ""
		} else {
			word += string(s[i])
		}
	}
	if word != "" {
		ts = append(ts, word)
	}

	for i := 1; i < len(ts); i++ {
		ts[i-1], ts[i] = ts[i], ts[i-1]
	}

	for i := 0; i < len(ts); i++ {
		for _, r := range ts[i] {
			z01.PrintRune(r)
		}
		if i < len(ts)-1 {
			z01.PrintRune(' ')
		} else {
			z01.PrintRune('\n')
		}
	}
}
