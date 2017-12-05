/**
 * @param {number} n
 * @return {boolean}
 */
var hasAlternatingBits = function(n) {
    var b = n.toString(2);
    var index = (b.indexOf("00") > b.indexOf("11")) ? b.indexOf("00") : b.indexOf("11");
    return index >= 0 ? false : true;
}