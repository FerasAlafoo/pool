package pool

import (
	"strings"
)

func WordFlip(str string) string {
	if str == "" {
		return "Invalid Output\n"
	}

	end := -1
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] != ' ' {
			end = i
			break
		}
	}
	if end == -1 {
		return "\n"
	}
	temp := strings.Split(str, " ")

	words := []string{}

	for _, r := range temp {
		if r != "" {
			words = append(words, r)
		}
	}
	toReturn := ""
	for i := len(words) - 1; i >= 0; i-- {
		toReturn += words[i]
		if i > 0 {
			toReturn += " "
		}
	}
	return toReturn + "\n"
}
