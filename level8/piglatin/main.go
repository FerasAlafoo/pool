package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]
	if arg == "" {
		return
	}
	v := false
	word := ""
	if vowel(string(arg[0])) {
		word += string(arg) + "ay"
		pHelp(word)
		return
	}
	for i := 1; i < len(os.Args[1]); i++ {
		if vowel(string(arg[i])) {
			v = true
			word = string(arg[i:len(os.Args[1])]) + string(arg[:i]) + "ay"
			break
		}
	}
	if !v {
		pHelp("No vowels")
	} else {
		pHelp(word)
	}
}

func vowel(r string) bool {
	if r == "a" || r == "e" || r == "i" || r == "o" || r == "u" || r == "A" || r == "E" || r == "I" || r == "O" || r == "U" {
		return true
	}
	return false
}

func pHelp(x string) {
	for _, r := range x {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
