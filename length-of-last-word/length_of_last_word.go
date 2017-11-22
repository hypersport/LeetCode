func lengthOfLastWord(s string) int {
    rl := 0
    sl := len(s)
    for i := sl - 1; i >= 0; i -- {
        if s[i] == ' ' && rl == 0 {
            continue
        } else if s[i] == ' ' {
            return rl
        } else {
            rl += 1
        }
    }
    return rl
}