func maxSubArray(nums []int) int {
    sum := nums[0]
    result := nums[0]
    for k, v := range nums {
        if k == 0 {
            continue
        }
        if sum < 0 {
            sum = v
        } else {
            sum += v
        }
        if result < sum {
            result = sum
        }
    }
    return result
}