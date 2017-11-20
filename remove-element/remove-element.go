func removeElement(nums []int, val int) int {
    var n = 0
    for _, v := range nums {
        if v != val {
            v, nums[n] = nums[n], v
            n ++
        }
    }
    return n
}
