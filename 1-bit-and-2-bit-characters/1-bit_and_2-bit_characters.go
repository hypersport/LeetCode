func isOneBitCharacter(bits []int) bool {
    var i int
    for i < len(bits) {
        if i == len(bits) - 1 {
            return true
        } else if bits[i] == 1 {
            i += 2
        } else {
            i += 1
        }
    }
    return false
}