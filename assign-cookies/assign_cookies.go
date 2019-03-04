func findContentChildren(g []int, s []int) int {
	gLen := len(g) - 1
	sLen := len(s) - 1
	quickSort(g, 0, gLen)
	quickSort(s, 0, sLen)
	result := 0
	for min(gLen, sLen) >= 0 {
		if s[sLen] >= g[gLen] {
			result++
			sLen--
		}
		gLen--
	}
	return result
}

func quickSort(a []int, left, right int) {
	if left < right {
		m := partition(a, left, right)
		quickSort(a, left, m-1)
		quickSort(a, m+1, right)
	}
}

func partition(a []int, low, high int) int {
	key := a[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if a[j] < key {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[high] = a[high], a[i+1]
	return i + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
