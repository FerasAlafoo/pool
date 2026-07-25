package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		return
	}

	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		valid := true
		s := Stack{}
		for _, r := range args[i] {
			switch r {
			case '(':
				s.Push(r)
			case '[':
				s.Push(r)
			case '{':
				s.Push(r)
			}

			switch r {
			case ')':
				top, b := s.Pop()
				if top != '(' || !b {
					valid = false
				}
			case ']':
				top, b := s.Pop()
				if top != '[' || !b {
					valid = false
				}
			case '}':
				top, b := s.Pop()
				if top != '{' || !b {
					valid = false
				}
			}
			if !valid {
				break
			}
		}

		if s.Size() == 0 && valid {
			fmt.Println("OK")
		} else {
			fmt.Print("Error")
		}
	}
}

type Stack struct {
	items []rune
}

func (s *Stack) Push(r rune) {
	s.items = append(s.items, r)
}

func (s *Stack) Pop() (rune, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	lastItem := len(s.items) - 1
	temp := s.items[lastItem]
	s.items = s.items[:lastItem]
	return temp, true
}

func (s *Stack) Size() int {
	return len(s.items)
}
