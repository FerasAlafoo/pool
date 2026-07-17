// package main

// import (
// 	"fmt"
// 	"pool/level5"
// )

// func main() {
// 	input1 := []uint{2, 3, 1, 1, 4}
// 	fmt.Println(pool.CanJump(input1))
// 	input2 := []uint{3, 2, 1, 0, 4}
// 	fmt.Println(pool.CanJump(input2))
// 	input3 := []uint{0}
// 	fmt.Println(pool.CanJump(input3))
// }

package main

import (
	"fmt"

	"pool/level5"
)

func main() {
	fmt.Println(pool.ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(pool.ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(pool.ConcatSlice([]int{1, 2, 3}, []int{}))
}
