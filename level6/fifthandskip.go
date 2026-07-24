// package piscine

// func FifthAndSkip(str string) string {
// 	x := ""
// 	group := ""
// 	skip := false

// 	if str == "" {
// 		return "\n"
// 	}
// 	if len(str) < 5 {
// 		return "Invalid Input\n"
// 	}

// 	for i := 0; i < len(str); i++ {
// 		if skip {
// 			skip = false
// 			continue
// 		}
// 		if str[i] != ' ' {
// 			group += string(str[i])
// 		}
// 		if len(group) == 5 {
// 			if len(x) != 0 {
// 				x += " "
// 			}
// 			x += group + " "
// 			group = ""
// 			skip = true
// 		}
// 	}

// 	if group != "" {
// 		if len(x) != 0 {
// 			x += " "
// 		}
// 		x += group
// 	}
// 	return x + "\n"
// }
/*
INSTRUCTIONS
Write a function FifthAndSkip() that takes a string and returns another string. The function separates every five
characters of the string with a space and removes the sixth one.
If there are spaces in the middle of a word, ignore them and get the first character after the spaces until you reach a
length of 5.
If the string is less than 5 characters return Invalid Input followed by a newline .
If the string is empty return a newline .
EXPECTED FUNCTION
USAGE
Here is a possible program to test your function:
OUTPUT
func FifthAndSkip(str string) string {
}
package main
import (
	"fmt	"
	"piscine	"
)
func main() {
fmt.Print(piscine.FifthAndSkip(	"abcdefghijklmnopqrstuwxyz	"))
fmt.Print(piscine.FifthAndSkip(	"This is a short sentence	"))
fmt.Print(piscine.FifthAndSkip(	"1234	"))
}
$ go run . | cat -e
abcde ghijk mnopq stuwx z$
Thisi ashor sente ce$
Invalid Input$
*/

package pool

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}
	countChars := 0
	for i := 0; i < len(str); i++ {
		if string(str[i]) != " " && string(str[i]) != "\t" {
			countChars++
		}
	}
	if countChars < 5 {
		return "Invalid Input\n"
	}
	x := ""
	count := 0
	for i := 0; i < len(str); i++ {
		if string(str[i]) != " " && string(str[i]) != "\t" {
			if count < 5 {
				x += string(str[i])
				count++
			} else {
				count = 0
				x += " "
			}
		}
	}

	return x + "\n"
}
