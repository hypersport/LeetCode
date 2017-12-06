func countBinarySubstrings(s string) int {
    var starti, lasti, sum int
    for i := 1; i < len(s); i ++ {
        if s[i] != s[i-1] {
            sum += 1
            lasti = i - starti
            starti = i
        } else if i - starti < lasti {
            sum += 1
        }
    }
    return sum
}