func isValid(s string) bool {
	var parentheses []rune
	var currentBracket rune

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			parentheses = append(parentheses, c)
		} else {
			if len(parentheses) > 0 {
                // pop operation
				currentBracket, parentheses = parentheses[len(parentheses)-1], parentheses[:len(parentheses)-1]
				if c == ')' && currentBracket != '(' {
					return false
				} else if c == ']' && currentBracket != '[' {
					return false
				} else if c == '}' && currentBracket != '{' {
					return false
				}
			} else {
				return false
			}
		}
	}

	if len(parentheses) != 0 {
		return false
	}
	return true
}
