package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg := os.Args[1]
	n, err := strconv.Atoi(arg)
	if err != nil {
		return
	}
	if n < 2 {
		return
	}

	for i := 2; i <= n; i++ {
		for n%i == 0 {
			if isPrime(i) {
				toPrint := i
				x := ""
				for toPrint > 0 {
					x = string(rune(toPrint%10+'0')) + x
					toPrint /= 10
				}
				for i := 0; i < len(x); i++ {
					z01.PrintRune(rune(x[i]))
				}
				if i < n {
					z01.PrintRune('*')
				}
				n /= i
			}
		}
	}
	z01.PrintRune('\n')
}

func isPrime(n int) bool {
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
