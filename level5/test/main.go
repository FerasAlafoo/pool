package main

import (
	"fmt"

	pool "pool/level5"
)

func main() {
	fmt.Println(pool.Chunk([]int{}, 10))
	fmt.Println(pool.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0))
	fmt.Println(pool.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3))
	fmt.Println(pool.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5))
	fmt.Println(pool.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4))
}

// package main

// import (
// 	"fmt"

// 	"pool/level5"
// )

// func main() {
// 	fmt.Println(pool.ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
// 	fmt.Println(pool.ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
// 	fmt.Println(pool.ConcatSlice([]int{1, 2, 3}, []int{}))
// }
