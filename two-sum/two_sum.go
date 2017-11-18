func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for k, v := range nums {
        if _, ok := m[target-v]; ok {
            if m[target-v] != k {
                return []int{m[target-v], k}
            }
        } else {
            m[v] = k
        }
    }
    return []int{}
}
