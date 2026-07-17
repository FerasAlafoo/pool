/*
Union
INSTRUCTIONS
Write a program that takes two string and displays, without doubles, the characters that appear in either one of the
string.
The display will be in the same order that the characters appear on the command line and will be followed by a
newline ( ' ' ).
If the number of arguments is different from 2, then the program displays a newline ( ' ' ).
USAGE
$ go run . zpadinton paqefwtdjetyiytjneytjoeyjnejeyj | cat -e
zpadintoqefwjy$
$
$ go run . ddf6vewg64f gtwthgdwthdwfteewhrtag6h4ffdhsd | cat -e
df6vewg4thras$
$
$ go run . rien "cette phrase ne cache rien" | cat -e
rienct phas$
$
$ go run . | cat -e
$
$ go run . rien | cat -e
$
*/
// take 2 strings, put together to become one string
// take first char in the string,
// if its not in the final yet, add it
// go to next char, see if its in the final, if not, add it, otherwise continue.

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Union() {
	if len(os.Args) != 3 {
		return
	}
	args := os.Args[1:]

	starter := []rune(args[0] + args[1])
	final := []rune{}
	for i := 0; i < len(starter); i++ {
		if len(final) == 0 {
			final = append(final, rune(starter[0]))
		}
		add := false
		index := 0
		for j := 0; j < len(final); j++ {
			if starter[i] == final[j] {
				add = false
				break
			} else {
				add = true
				index = i
			}
		}
		if add {
			final = append(final, rune(starter[index]))
		}
	}

	for _, r := range final {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
