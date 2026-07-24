package main

import (
	"fmt"

	pool "pool/level6"
)

func main() {
	fmt.Print(pool.FifthAndSkip("abcdefghijklmnopqrstuwxyz	"))
	fmt.Print(pool.FifthAndSkip("This is a short sentence	"))
	fmt.Print(pool.FifthAndSkip("1234	"))
}
