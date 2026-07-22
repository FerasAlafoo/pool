package main

import (
	"fmt"
	"pool"
)
func main() {
	newStack := pool.Stack{}
	
	newStack.Push(10)
	newStack.Push(20)
	newStack.Push(30)

	fmt.Println(newStack)

	newStack.Pop()
	fmt.Println(newStack)

}
