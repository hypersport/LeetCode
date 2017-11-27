/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function(n) {
    if (n <= 0) {
        return 0;
    }
    var s = [0, 1];
    var tmp = 0;
    for (var i = 0; i < n; i ++) {
        tmp = s[0];
        s[0] = s[1];
        s[1] += tmp;
    }
    return s[1];
};