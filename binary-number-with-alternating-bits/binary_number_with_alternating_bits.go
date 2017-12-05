func hasAlternatingBits(n int) bool {
    remainder := n % 2
    n /= 2
    for n > 0 {
        new_remainder := n % 2
        if new_remainder == remainder {
            return false
        }
        n /= 2
        remainder = new_remainder
    }
    return true
}