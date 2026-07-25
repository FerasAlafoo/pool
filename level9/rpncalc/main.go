/*
INSTRUCTIONS
Write a program that takes a string which contains an equation written in Reverse Polish Notation (RPN) as its first
argument, that evaluates the equation, and that prints the result on the standard output followed by a newline ( '
' ).
Reverse Polish Notation is a mathematical notation in which every operator follows all of its operands. In RPN,
every operator encountered evaluates the previous 2 operands, and the result of this operation then becomes the
first of the two operands for the subsequent operator. Operands and operators must be spaced by at least one
space.
The following operators must be implemented: + , - , * , / , and % .
If the string is not valid or if there is not exactly one argument, Error must be printed on the standard output
followed by a newline. If the string has extra spaces it is still considered valid.
All the given operands must fit in an int .
FORMULA CONVERSION EXAMPLES
3 + 4 >> 3 4 +
((1 * 2) * 3) - 4 >> 1 2 * 3 * 4 - or 3 1 2 * * 4 -
50 * (5 - (10 / 9)) >> 5 10 9 / - 50 *
EVALUATION TRACE EXAMPLES
USAGE
•
•
•
1 2 * 3 * 4 -
2 3 * 4 -
6 4 -
2
Or:
3 1 2 * * 4 -
3 2 * 4 -
6 4 -
2
$ go run . "1 2 * 3 * 4 +" | cat -e
10$
$ go run . "1 2 3 4 +" | cat -e
Error$
$ go run . | cat -e
Error$
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}

	stuff := strings.Split(os.Args[1], " ")
	arg := []string{}
	for _, r := range stuff {
		if r != "" {
			arg = append(arg, r)
		}
	}

	newStack := Stack{}

	for i := 0; i < len(arg); i++ {
		isOp := false
		op := ""
		switch arg[i] {
		case "+", "-", "*", "/", "%":
			op = arg[i]
			isOp = true
		default:
			n, err := strconv.Atoi(arg[i])
			if err != nil {
				fmt.Println("Error")
				return
			}
			newStack.Push(n)
		}

		if isOp {
			if newStack.Size() < 2 {
				fmt.Println("Error")
				return
			}
			if op == "+" {
				newStack.Push(newStack.Pop() + newStack.Pop())
			}
			if op == "*" {
				newStack.Push(newStack.Pop() * newStack.Pop())
			}
			if op == "-" {
				b := newStack.Pop()
				a := newStack.Pop()
				newStack.Push(a - b)
			}
			if op == "/" {
				b := newStack.Pop()
				a := newStack.Pop()
				if b == 0 {
					fmt.Println("Error")
					return
				}
				newStack.Push(a / b)
			}
			if op == "%" {
				b := newStack.Pop()
				a := newStack.Pop()
				if b == 0 {
					fmt.Println("Error")
					return
				}
				newStack.Push(a % b)
			}
		}
	}
	if newStack.Size() != 1 {
		fmt.Println("Error")
		return
	}
	fmt.Println(newStack.Pop())
}

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	lastItem := len(s.items) - 1
	temp := s.items[lastItem]
	s.items = s.items[:lastItem]
	return temp
}

func (s *Stack) Size() int {
	return len(s.items)
}
