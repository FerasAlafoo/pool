/*
ConcatSlice
INSTRUCTIONS
Write a function ConcatSlice() that takes two slices of integers as arguments and returns the concatenation of
the two slices.
EXPECTED FUNCTION
USAGE
Here is a possible program to test your function:
EXPECTED OUTPUT
func ConcatSlice(slice1, slice2 []int) []int {
}
package main
import (
"fmt"
"piscine"
)
func main() {
fmt.Println(piscine.ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
fmt.Println(piscine.ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
fmt.Println(piscine.ConcatSlice([]int{1, 2, 3}, []int{}))
}
*/

package pool

func ConcatSlice(slice1, slice2 []int) []int {
	x := []int{}
	for i := 0; i < len(slice1); i++ {
		x = append(x, slice1[i])
	}
	for i := 0; i < len(slice2); i++ {
		x = append(x, slice2[i])
	}
	return x
}
