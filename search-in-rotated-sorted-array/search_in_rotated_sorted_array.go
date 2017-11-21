func search(nums []int, target int) int {
    for k, v := range nums {
        if v == target {
            return k
        }
    }
    return -1
}