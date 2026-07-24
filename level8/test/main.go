package main

import (
	"fmt"

	pool "pool/level8"
)

func main() {
	fmt.Println(pool.ItoaBase(10, 10))
	fmt.Println(pool.ItoaBase(10, 2))
	fmt.Println(pool.ItoaBase(10, 16))
	fmt.Println(pool.ItoaBase(0, 2))
	fmt.Println(pool.ItoaBase(-10, 10))
	fmt.Println(pool.ItoaBase(-255, 16))
	fmt.Println(pool.ItoaBase(123456, 16))
	fmt.Println(pool.ItoaBase(-123456789, 16))
}
