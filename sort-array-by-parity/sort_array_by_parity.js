/**
 * @param {number[]} A
 * @return {number[]}
 */
var sortArrayByParity = function(A) {
    return A.sort((i,j) => i % 2 - j % 2)
};