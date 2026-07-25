/*
INSTRUCTIONS
Write a program that prints in descending order on a single line all unique combinations of three different digits so
that the first digit is greater than the second and the second is greater than the third.
These combinations are separated by a comma and a space.
999 or 000 are not valid combinations because the digits are not different.
789 should not be shown because the first digit is not greater than the second.
USAGE
Here is an incomplete output :
$ go run . | cat -e
987, 986, 985, 984, 983, 982, 981, 980, 976, ..., 310, 210$
*/
package main

import "github.com/01-edu/z01"

func main() {
	for i := '9'; i >= '2'; i-- {
		for j := '8'; j >= '1'; j-- {
			for k := '7'; k >= '0'; k-- {
				if i > j && j > k {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					// if not at end, print comma then space, else new line
				}
			}
		}
	}
}
