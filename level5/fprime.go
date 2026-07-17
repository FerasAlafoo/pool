package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	fmt.Print(args[0])
	z01.PrintRune(1)
	z01.PrintRune('a')
}

func isPrime(n int, divider int) bool {
	if n < 2 {
		return false
	}
	for i := divider; i*i >= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
