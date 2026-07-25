package main

import (
	"fmt"

	pool "pool/level4"
)

func main() {
	fmt.Print(pool.FromTo(1, 10))
	fmt.Print(pool.FromTo(10, 1))
	fmt.Print(pool.FromTo(10, 10))
	fmt.Print(pool.FromTo(100, 10))
}
