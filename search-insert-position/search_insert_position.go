func searchInsert(nums []int, target int) int {
    for k, v := range nums {
        if v > target {
            return 0
        } else if v == target {
            return k
        } else if (k == len(nums) - 1) || (v < target && nums[k+1] > target) {
            return k + 1
        }
    }
    return 0
}