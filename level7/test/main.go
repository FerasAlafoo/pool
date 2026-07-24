package main

import (
	"fmt"

	pool "pool/level7"
)

func main() {
	fmt.Print(pool.WordFlip("First second last"))
	fmt.Print(pool.WordFlip(""))
	fmt.Print(pool.WordFlip(" "))
	fmt.Print(pool.WordFlip(" hello all of you! "))
}
