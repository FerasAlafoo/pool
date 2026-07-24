/*
INSTRUCTIONS
Write a function RevConcatAlternate() that receives two slices of int as arguments and returns a new slice with
alternated values of each slice in reverse order.
The input slices can have different lengths.
The new slice should start with the elements from the largest slice first and when they became equal size
slices, it should add an element of the first given slice.
If the slices are of equal length, the new slice should start with an element of the first slice.
Note: you can check the examples below for more details.
EXPECTED FUNCTION
USAGE
Here is a possible program to test your function:
OUTPUT
func RevConcatAlternate(slice1, slice2 []int) []int {
}
package main
import (

	"fmt	"
	"piscine	"

)
func main() {
fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{}))
}
$ go run .
[3 6 2 5 1 4]
[9 8 7 3 6 2 5 1 4]
[8 9 3 2 5 1 4]
[3 2 1]
*/
package pool

func RevConcatAlternate(slice1, slice2 []int) []int {
	first := slice1
	second := slice2

	if len(first) < len(second) {
		first, second = second, first
	}

	x := []int{}
	l1 := len(first) - 1
	l2 := len(second) - 1

	for l1 > l2 {
		x = append(x, first[l1])
		l1--
	}

	for i := 0; i < len(second); i++ {
		x = append(x, slice1[l1])
		x = append(x, slice2[l2])
		l1--
		l2--
	}

	return x
}
