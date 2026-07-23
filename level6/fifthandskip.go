package piscine

func FifthAndSkip(str string) string {
	x := ""
	group := ""
	skip := false

	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Input\n"
	}

	for i := 0; i < len(str); i++ {
		if skip {
			skip = false
			continue
		}
		if str[i] != ' ' {
			group += string(str[i])
		}
		if len(group) == 5 {
			if len(x) != 0 {
				x += " "
			}
			x += group + " "
			group = ""
			skip = true
		}
	}

	if group != "" {
		if len(x) != 0 {
			x += " "
		}
		x += group
	}
	return x + "\n"
}
