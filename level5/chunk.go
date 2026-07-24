/* Write a function called Chunk that receives as parameters a slice, slice []int , and a number size int . The
goal of this function is to chunk a slice into many sub slices where each sub slice has the length of size.
If the size is 0 it should print a newline ( ' ' ).
EXPECTED FUNCTION
USAGE
Here is a possible program to test your function :
EXPECTED OUTPUT
func Chunk(slice []int, size int) {
}
package main
func main() {
Chunk([]int{}, 10)
Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
$ go run .
[]
[[0 1 2] [3 4 5] [6 7]]
[[0 1 2 3 4] [5 6 7]]
[[0 1 2 3] [4 5 6 7]]
*/

package pool

func Chunk(slice []int, size int) [][]int {
	// step 1 : loop over the slice, appending to a slice called x

	c := 0
	toReturn := [][]int{}
	x := []int{}
	if len(slice) == 0 {
		return [][]int(nil)
	}
	if size == 0 {
		return nil
	}
	for i := 0; i < len(slice); i++ {
		if c < size {
			x = append(x, slice[i])
			c++
		} else {
			toReturn = append(toReturn, x)
			c = 0
			x = []int{}
			i--
		}
	}
	if len(x) != 0 {
		toReturn = append(toReturn, x)
	}

	return toReturn
}
