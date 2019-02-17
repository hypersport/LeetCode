func licenseKeyFormatting(S string, K int) string {
	runes := make([]rune, 0)
	for _, c := range S {
		if c != '-' {
			runes = append(runes, c)
		}
	}
	res := make([]rune, 0)
	i := 0
	for _, c := range reverse(runes) {
		if i != K {
			res = append(res, c)
			i++
		} else {
			res = append(res, '-')
			res = append(res, c)
			i = 1
		}
	}
	return reverse(toupper(res))
}

func reverse(runes []rune) string {
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func toupper(runes []rune) []rune {
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 97 && runes[i] <= 122 {
			runes[i] -= 32
		}
	}
	return runes
}
