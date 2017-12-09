/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isSubsequence = function(s, t) {
    var index = 0;
    for (var i = 0; i < s.length; i ++) {
        index = t.indexOf(s.charAt(i), index);
        if (index < 0) {
            return false;
        }
        index += 1;
    }
    return true;
};