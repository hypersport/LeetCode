func sortArrayByParity(A []int) []int {
	left, right := 0, len(A)-1
	for left < right {
		if A[left]%2 == 1 && A[right]%2 == 0 {
			A[left], A[right] = A[right], A[left]
			left++
			right--
		} else {
			if A[left]%2 == 0 {
				left++
			}
			if A[right]%2 == 1 {
				right--
			}
		}
	}
	return A
}
