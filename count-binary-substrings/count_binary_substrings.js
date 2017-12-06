/**
 * @param {string} s
 * @return {number}
 */
var countBinarySubstrings = function(s) {
    var starti = 0, lasti = 0, sum = 0;
    for (var i = 1; i < s.length; i ++) {
        if (s.charAt(i) != s.charAt(i-1)) {
            sum += 1;
            lasti = i - starti;
            starti = i;
        } else if (i - starti < lasti) {
            sum += 1;
        }
    }
    return sum;
};