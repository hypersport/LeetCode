/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
    if (x<0 || (x!=0 && x%10==0)) {
        return false;
    }
    var y = 0;
    while (x>y) {
        y = y*10 + x%10;
        x = parseInt(x/10);
    }
    return (x==y || x==parseInt(y/10));
};