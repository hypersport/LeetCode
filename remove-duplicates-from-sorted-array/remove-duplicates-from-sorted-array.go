func removeDuplicates(nums []int) int {
    var l := 0
    for k, v := range nums {
        if k == 0 || v != nums[k-1] {
            nums[l] = v
            l += 1
        }
    }
    return l
}
