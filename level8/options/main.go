package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}

	args := os.Args[1:]
	status := ""
	for i := 0; i < len(args); i++ {
		a := args[i]
		if len(a) > 1 && a[0] == '-' && a[1] == 'h' {
			status = "options"
			break
		}
		if len(a) == 1 {
			status = "invalid"
			break
		}
		for j := 1; j < len(a); j++ {
			if a[0] != '-' || (a[j] < 'a' || a[j] > 'z') {
				status = "invalid"
			}
		}
		if status == "invalid" {
			break
		}
	}

	if status == "options" {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}
	if status == "invalid" {
		fmt.Println("Invalid Option")
		return
	}

	flags := 0
	for _, a := range args {
		for i := 1; i < len(a); i++ {
			index := a[i] - 'a'
			flags |= 1 << index
		}
	}

	for i := 31; i >= 0; i-- {
		if flags&(1<<i) != 0 {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}

		if i%8 == 0 && i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
