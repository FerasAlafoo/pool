// package main

// import (
// 	"os"
// )

// func main() {
// 	for i := 1; i < len(os.Args) ; i++ {
// 		result := ""
// 		sent := os.Args[i]

// 		for j := 0 ; j <len(sent) ; j++ {
// 			letter := sent[j]

// 			if alphabetic {
// 				if upper {
// 					letter += 32
// 				}
// 				if j+1 == len(sent) || !alphabetic {
// 					letter -= 32
// 				}
// 			}
// 			result += string(letter)
// 		}
// 	}
// }
/*
INSTRUCTIONS
Write a program that takes one or more arguments and that, for each argument, puts the last letter of each word in
uppercase and the rest in lowercase. It displays the result followed by a newline ( ' ' ).
If there are no argument, the program displays nothing.
USAGE
$ go run . "First SMALL TesT" | cat -e
firsT smalL tesT$
$ go run . "SEconD Test IS a LItTLE EasIEr" "bEwaRe IT'S NoT HARd WhEN " " Go a dernier
0123456789 for the road e" | cat -e
seconD tesT iS A littlE easieR$
bewarE it'S noT harD wheN $
 gO A dernieR 0123456789 foR thE roaD E$
$ go run .
$
*/

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		x := []rune(args[i])
		for j := 0; j < len(x); j++ {
			if j == len(x)-1 && x[j] != ' ' {
				if lower(x[j]) {
					x[j] -= 32
				}
			} else if x[j] != ' ' && x[j+1] == ' ' {
				if lower(x[j]) {
					x[j] -= 32
				}
			} else if upper(x[j]) {
				x[j] += 32
			}
		}
		for _, r := range x {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

func lower(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func upper(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}
