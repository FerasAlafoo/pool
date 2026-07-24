/*
INSTRUCTIONS
The function receives a slice of strings and one or more integers, and returns a slice of strings. The returned slice is
part of the received one but cut from the position indicated in the first int, until the position indicated by the second
int.
In case there only exists one int, the resulting slice begins in the position indicated by the int and ends at the end of
the received slice.
The integers can be negative.
EXPECTED FUNCTION
USAGE
Here is a possible program to test your function :
func Slice(a []string, nbrs... int) []string {
}
package main
import (

	"fmt	"
	"piscine	"

)

	func main() {
	 a := []string{	"coding	", 	"algorithm	", 	"ascii	", 	"package	", 	"golang	"}
	 fmt.Printf(	"%#v
		", piscine.Slice(a, 1))
	 fmt.Printf(	"%#v
		", piscine.Slice(a, 2, 4))
	 fmt.Printf(	"%#v
		", piscine.Slice(a, -3))
	 fmt.Printf(	"%#v
		", piscine.Slice(a, -2, -1))
	 fmt.Printf(	"%#v
		", piscine.Slice(a, 2, 0))
	}

$ go run .
[]string{	"algorithm	", 	"ascii	", 	"package	", 	"golang	"}
[]string{	"ascii	", 	"package	"}
[]string{	"ascii	", 	"package	", 	"golang	"}
[]string{	"package	"}
[]string(nil)
*/
package pool

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 {
		return a
	}

	var x []string
	l := len(a)
	if len(nbrs) == 2 {
		if nbrs[1] < 0 {
			nbrs[1] += len(a)
		}
		l = nbrs[1]
	}
	if nbrs[0] < 0 {
		nbrs[0] += len(a)
	}
	for i := nbrs[0]; i < l; i++ {
		x = append(x, a[i])
	}

	return x
}
