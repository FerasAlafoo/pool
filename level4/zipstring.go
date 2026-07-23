package piscine

import "strconv"

func ZipString(s string) string {
	if s == "" {
		return s
	}

	result := ""
	count := 1

	for i := 1; i <= len(s); i++ {
		if i < len(s) && s[i] == s[i-1] {
			count++
		} else {
			result += strconv.Itoa(count) + string(s[i-1])
			count = 1
		}
	}
	return result
	}