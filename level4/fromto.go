/*
INSTRUCTIONS
Write a function that takes two integers and returns a string showing the range of numbers from the first to the
second.
The numbers must be separated by a comma and a space.
If any of the arguments is bigger than 99 or less than 0, the function returns Invalid followed by a newline
\n .
Prepend a 0 to any number that is less than 10.
Add a new line \n at the end of the string.
EXPECTED FUNCTION
USAGE
Here is a possible program to test your function:
EXPECTED OUTPUT
•
•
•
•
func FromTo(from int, to int) string {
}
package main
import (

	"fmt	"
	"piscine	"

)
func main() {
fmt.Print(piscine.FromTo(1, 10))
fmt.Print(piscine.FromTo(10, 1))
fmt.Print(piscine.FromTo(10, 10))
fmt.Print(piscine.FromTo(100, 10))
}
$ go run . | cat -e
01, 02, 03, 04, 05, 06, 07, 08, 09, 10$
10, 09, 08, 07, 06, 05, 04, 03, 02, 01$
10$
Invalid$
*/
package pool

import "strconv"

func FromTo(from int, to int) string {
	if (from < 0 || from > 99) || (to < 0 || to > 99) {
		return "Invalid\n"
	}
	if from == to {
		return strconv.Itoa(from) + "\n"
	}
	x := ""

	if from < to {
		for i := from; i < to; i++ {
			if i < 10 {
				x += "0"
			}
			x += strconv.Itoa(i) + ", "
		}
		x += strconv.Itoa(to)
	}
	if from > to {
		for i := from; i > to; i-- {
			if i < 10 {
				x += "0"
			}
			x += strconv.Itoa(i) + ", "
		}
		if to < 10 {
			x += "0"
		}
		x += strconv.Itoa(to)
	}
	return x + "\n"
}
