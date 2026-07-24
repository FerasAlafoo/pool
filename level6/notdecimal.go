/*
INSTRUCTIONS
Write a function called NotDecimal() that takes as an argument a string in form of a float number with the
decimal point and returns a string converted to int without the decimal point (you will have to multiply it by 10^n to
remove the .).
If the number doesn't have a decimal point or there is only a zero after the . return the number followed by a
newline .
If the argument is empty return a newline .
If the argument is not a number return it followed by a newline .
EXPECTED FUNCTION
USAGE
Here is a possible program to test your function:
OUTPUT
func NotDecimal(dec string) string {
}
package main
import (
	"fmt	"
)
func main() {
fmt.Print(NotDecimal(	"0.1	"))
fmt.Print(NotDecimal(	"174.2	"))
fmt.Print(NotDecimal(	"0.1255	"))
fmt.Print(NotDecimal(	"1.20525856	"))
fmt.Print(NotDecimal(	"-0.0f00d00	"))
fmt.Print(NotDecimal(	"	"))
fmt.Print(NotDecimal(	"-19.525856	"))
fmt.Print(NotDecimal(	"1952	"))
}
$ go run . | cat -e
1$
1742$
1255$
120525856$
-0.0f00d00$
$
-19525856$
1952$
*/

package pool

func NotDecimal(dec string) string {
	if dec == "" {
		return dec
	}

	dotIndex := -1
	dot := false
	for i := 0; i < len(dec); i++ {
		if dec[i] == '.' {
			if dot {
				return dec
			}
			dotIndex = i
			dot = true
		}
	}
	if dotIndex == -1 {
		return dec
	}
	// SEE THIS PART BELOW
	for i := 0; i < len(dec); i++ {
		if !(dec[i] == '-' && i == 0) && (dec[i] < '0' || dec[i] > '9') && (dec[i] != '.') && (dec[i] != ' ') && (dec[i] != '\t') {
			return dec
		}
	}
	dec = dec[:dotIndex] + dec[dotIndex+1:]
	sign := ""
	if dec[0] == '-' {
		sign = "-"
		dec = dec[1:]
	}
	start := 0
	for i := 0; i < len(dec); i++ {
		if dec[i] >= '1' && dec[i] <= '9' {
			start = i
			break
		}
	}
	endIndex := len(dec)
	for i := len(dec) - 1; i > dotIndex; i-- {
		if dec[i] >= '1' && dec[i] <= '9' {
			endIndex = i + 1
			break
		}
	}
	x := sign
	for i := start; i < endIndex; i++ {
		x += string(dec[i])
	}

	return x
}
