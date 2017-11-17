func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for k, v := range nums {
        m[v] = k
    }
    for k, v := range nums {
        if _, ok := m[target-v]; ok {
            if m[target-v] != k {
                return []int{k, m[target-v]}
            }
        }
    }
    return []int{}
}
