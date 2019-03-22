func shortestToChar(S string, C byte) []int {
	indexes := make([]int, 0)
	for i, c := range S {
		if byte(c) == C {
			indexes = append(indexes, i)
		}
	}

	result := make([]int, len(S))
	index, ind, stop := 0, 0, 0
	for index = 1; index < len(indexes); index++ {
		for ind = stop; ind < len(S); ind++ {
			if ind == indexes[index] {
				result[ind] = 0
				stop = ind + 1
				break
			} else if ind <= indexes[index-1] {
				result[ind] = indexes[index-1] - ind
			} else if ind > indexes[index-1] && ind < indexes[index] {
				result[ind] = min(ind-indexes[index-1], indexes[index]-ind)
			}
		}
	}

	for len(S) > ind {
		result[ind] = max(ind-indexes[index-1], indexes[index-1]-ind)
		ind++
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
