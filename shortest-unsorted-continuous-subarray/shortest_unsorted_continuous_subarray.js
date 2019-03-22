/**
 * @param {number[]} nums
 * @return {number}
 */
var findUnsortedSubarray = function(nums) {
    var [ ...ori_nums ] = nums;
    var indexes = [];
    nums.sort(sortNum);
    for (var i = 0; i < nums.length; i ++) {
        if (ori_nums[i] != nums[i]) {
            indexes.push(i);
        }
    }
    if (indexes.length === 0) {
        return 0;
    }
    return indexes[indexes.length - 1] - indexes[0] + 1
};

function sortNum(a, b) {
    return a - b;
}