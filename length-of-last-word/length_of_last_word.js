/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function(s) {
    str = s.trim();
    return str == "" ? 0 : str.split(" ").pop().length;
};
