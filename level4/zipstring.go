/*
INSTRUCTIONS
Write a function that takes a string and returns a new string that replaces every character with the number of
duplicates and the character itself, deleting the extra duplications.
The letters are from the latin alphabet list only. Any other character, symbols, shall not be tested.
EXPECTED FUNCTION
USAGE
Here is a possible program to test your function:
EXPECTED OUTPUT
func ZipString(s string) string {
}
package main
import (

	"fmt	"

)
func main() {
fmt.Println(ZipString(	"YouuungFellllas	"))
fmt.Println(ZipString(	"Thee quuick browwn fox juumps over the laaazy dog	"))
fmt.Println(ZipString(	"Helloo Therre!	"))
}
$ go run .
1Y1o3u1n1g1F1e4l1a1s
1T1h2e1 1q2u1i1c1k1 1b1r1o2w1n1 1f1o1x1 1j2u1m1p1s1 1o1v1e1r1 1t1h1e1 1l3a1z1y1 1d1o1g
1H1e2l2o1 1T1h1e2r1e1!
*/
package pool

import (
	"strconv"
)

func ZipString(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		c := 1
		char := s[i]
		for j := i + 1; j < len(s); j++ {
			if char == s[j] {
				c++
				i++
			} else {
				break
			}
		}
		result += strconv.Itoa(c) + string(char)
	}
	if result[len(result)-1] == ' ' {
		return result[:len(result)-3]
	}
	return result
}
