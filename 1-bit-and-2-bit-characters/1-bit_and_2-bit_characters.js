/**
 * @param {number[]} bits
 * @return {boolean}
 */
var isOneBitCharacter = function(bits) {
    var i = 0;
    while (i < bits.length) {
        if (i == bits.length - 1) {
            return true;
        } else if (bits[i] == 1) {
            i += 2;
        } else {
            i += 1;
        }
    }
    return false;
};