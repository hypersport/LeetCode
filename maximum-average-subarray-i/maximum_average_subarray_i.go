func findMaxAverage(nums []int, k int) float64 {
	currentSum, maxKSum := nums[0], nums[0]
	if k == 1 {
		for _, n := range nums {
			if maxKSum < n {
				maxKSum = n
			}
		}
	} else {
		for _, n := range nums[1:k] {
			currentSum += n
		}
		maxKSum = currentSum
		for i := 0; i < len(nums)-k; i++ {
			currentSum = currentSum - nums[i] + nums[i+k]
			if maxKSum < currentSum {
				maxKSum = currentSum
			}
		}
	}
	return float64(maxKSum) / float64(k)
}
