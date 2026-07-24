/*
INSTRUCTIONS
Write a program named hiddenp that takes two strings as arguments. The program should check if the first string
s1 is hidden in the second s2 . s1 is considered hidden in s2 if it is possible to find each character from s1 in
s2 , in the same order as they appear in s1 , but not necessarily consecutively.
If s1 is hidden in s2 , the program should display 1 followed by a newline.
If s1 is not hidden in s2 , the program should display 0 followed by a newline.
If s1 is an empty string, it is considered hidden in any string.
If the number of arguments is different from 2, the program should display nothing.
USAGE
•
•
•
•
$ go run . "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6" | cat -e
1$
$ go run . "abc" "2altrb53c.sse" | cat -e
1$
$ go run . "abc" "btarc" | cat -e
0$
$ go run . "DD" "DABC" | cat -e
0$
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

	if os.Args[1] == "" {
		z01.PrintRune('1')
		z01.PrintRune('\n')
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]
	start := 0

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
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}
	}
	z01.PrintRune('1')
	z01.PrintRune('\n')
}
