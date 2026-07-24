package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	num := 0
	for _, r := range os.Args[1] {
		if r < '0' || r > '9' {
			fmt.Printf("ERROR: cannot convert to roman digit\n")
			return
		}
		num = num*10 + int(r-'0')
	}

	if num <= 0 || num >= 4000 {
		fmt.Printf("ERROR: cannot convert to roman digit\n")
		return
	}

	values := []int{
		1000,
		900,
		500,
		400,
		100,
		90,
		50,
		40,
		10,
		9,
		5,
		4,
		1,
	}

	symbols := []string{
		"M",
		"(M-C)",
		"D",
		"(D-C)",
		"C",
		"(C-X)",
		"L",
		"(L-X)",
		"X",
		"(X-I)",
		"V",
		"(V-I)",
		"I",
	}

	roman := []string{
		"M",
		"CM",
		"D",
		"CD",
		"C",
		"XC",
		"L",
		"XL",
		"X",
		"IX",
		"V",
		"IV",
		"I",
	}

	calc := ""
	result := ""

	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			num -= values[i]
			calc += symbols[i]
			result += roman[i]

			if num != 0 {
				calc += "+"
			}
		}
	}

	fmt.Printf("%s\n", calc)
	fmt.Printf("%s\n", result)
}
