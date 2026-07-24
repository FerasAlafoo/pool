/*
INSTRUCTIONS
Write a program that takes two string and checks whether it is possible to write the first string with characters from
the second string. This rewrite must respect the order in which these characters appear in the second string.
If it is possible, the program displays the string followed by a newline ( ' ' ), otherwise it simply displays nothing.
If the number of arguments is different from 2, the program displays nothing.
USAGE
$ go run . 123 123
123
$ go run . faya fgvvfdxcacpolhyghbreda
faya
$ go run . faya fgvvfdxcacpolhyghbred
$ go run . error rrerrrfiiljdfxjyuifrrvcoojh
$ go run . "quarante deux" "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq "
quarante deux
$ go run .
$
*/

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	start := 0
	s1 := os.Args[1]
	s2 := os.Args[2]
	for i := 0; i < len(s1); i++ {
		found := false
		for j := start; j < len(s2); j++ {
			if s1[i] == s2[j] {
				found = true
				start = j + 1
				break
			}
		}
		if !found {
			return
		}
	}

	for _, r := range s1 {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
