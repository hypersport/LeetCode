func lengthOfLIS(nums []int) int {
	numsLen := len(nums)
	if numsLen <= 0 {
		return 0
	}
	result := 1
	var dp []int = make([]int, numsLen, numsLen)
	for i := 0; i < numsLen; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				result = max(result, dp[i])
			}
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
