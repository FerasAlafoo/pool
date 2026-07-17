/*
AddPrimeSum
INSTRUCTIONS
Write a program that takes a positive integer as argument and displays the sum of all prime numbers inferior or
equal to it followed by a newline (' ').
If the number of arguments is different from 1, or if the argument is not a positive number, the program displays 0
followed by a newline.
USAGE
$ go run . 5
10
$ go run . 7
17
$ go run . -2
0
$ go run . 0
0
$ go run .
0
$ go run . 5 7
0
*/
package pool

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func AddPrimeSum() {
	if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune(10)
		return
	}

	n, err := strconv.Atoi(os.Args[1])

	if err != nil || n < 2 {
		z01.PrintRune('0')
		z01.PrintRune(10)
		return
	}

	sum := 0
	for i := 0; i <= n; i++ {
		if prime(i) {
			sum += i
		}
	}

	print := ""
	for sum > 0 {
		print = string(rune(sum%10+'0')) + print
		sum /= 10
	}
	for _, r := range print {
		z01.PrintRune(r)
	}
	if print != "" {
		z01.PrintRune('\n')
	}
}

func prime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
