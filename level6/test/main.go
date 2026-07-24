// package main

// import (
// 	"fmt"

// 	pool "pool/level6"
// )

// func main() {
// 	fmt.Print(pool.FifthAndSkip("abcdefghijklmnopqrstuwxyz	"))
// 	fmt.Print(pool.FifthAndSkip("This is a short sentence	"))
// 	fmt.Print(pool.FifthAndSkip("1234	"))
// }

// package main

// import (
// 	"fmt"

// 	pool "pool/level6"
// )

// func main() {
// 	fmt.Println(pool.RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
// 	fmt.Println(pool.RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
// 	fmt.Println(pool.RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
// 	fmt.Println(pool.RevConcatAlternate([]int{1, 2, 3}, []int{}))
// }

// package main

// import (
// 	"fmt"

// 	pool "pool/level6"
// )

// func main() {
// 	a := []string{"coding	", "algorithm	", "ascii	", "package	", "golang	"}
// 	fmt.Printf("%#v", pool.Slice(a, 1))
// 	fmt.Printf("%#v", pool.Slice(a, 2, 4))
// 	fmt.Printf("%#v", pool.Slice(a, -3))
// 	fmt.Printf("%#v", pool.Slice(a, -2, -1))
// 	fmt.Printf("%#v", pool.Slice(a, 2, 0))
// }

package main

import (
	"fmt"

	pool "pool/level6"
)

func main() {
	fmt.Println(pool.NotDecimal("0.1	"))
	fmt.Println(pool.NotDecimal("174.2	"))
	fmt.Println(pool.NotDecimal("0.1255	"))
	fmt.Println(pool.NotDecimal("1.20525856	"))
	fmt.Println(pool.NotDecimal("-0.0f00d00	"))
	fmt.Println(pool.NotDecimal("	"))
	fmt.Println(pool.NotDecimal("-19.525856	"))
	fmt.Println(pool.NotDecimal("1952	"))
}
