/*
INSTRUCTIONS
Write a function that takes two strings and returns the number of characters that are not included in both, without
repeating characters.
If there are no unique characters return 0.
If both strings are empty return -1.
EXPECTED FUNCTION
USAGE
Here is a possible program to test your function:
EXPECTED OUTPUT
•
•
func WeAreUnique(str1 , str2 string) int {
}
package main
import (
	"fmt	"
)
func main() {
fmt.Println(WeAreUnique(	"foo	", 	"boo	"))
fmt.Println(WeAreUnique(	"	", 	"	"))
fmt.Println(WeAreUnique(	"abc	", 	"def	"))
}
$ go run .
2
-1
6
*/

package pool

import (
	"strings"
)

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	c := 0
	seen := ""

	for i := 0; i < len(str1); i++ {
		char := string(str1[i])
		if !strings.Contains(str2, char) && !strings.Contains(seen, char) {
			seen += char
			c++
		}
	}

	for i := 0; i < len(str2); i++ {
		char := string(str2[i])
		if !strings.Contains(str1, char) && !strings.Contains(seen, char) {
			seen += char
			c++
		}
	}
	return c
}
