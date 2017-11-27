func climbStairs(n int) int {
    if n <=0 {
        return 0
    }
    var s = []int{0, 1}
    for i := 0; i < n; i++ {
        s[0], s[1] = s[1], s[0] + s[1]
    }
    return s[1]
}