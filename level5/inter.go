/*
Inter
INSTRUCTIONS
Write a program that takes two string and displays, without doubles, the characters that appear in both string, in the
order they appear in the first one.
The display will be followed by a newline (' ').
If the number of arguments is different from 2, the program displays nothing.
USAGE
$ go run . "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
padinto
$ go run . ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd
df6ewg4
*/
package pool

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	args := os.Args[1:]
	arg1 := []rune(args[0])
	arg2 := []rune(args[1])
	output := ""

	for i := 0 ; i < len(arg1) ; i++ {
		l1 := arg1[i]
		for j := 0 ; j < len(arg2) ; j++ {
			l2 := arg2[j]
			add := false
			for k := 0 ; k < len(output) ; k++ {
				if l1 == l2 && (len(output) != 1 || l1 != rune(output[k])) {
					add = true
				} else {
					add = false
				}
			}
			if add {
				output += string(l2)
			}
		}
	}
	fmt.Print(output)
}
