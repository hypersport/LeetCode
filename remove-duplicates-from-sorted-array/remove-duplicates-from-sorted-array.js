/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
    var l = 0;
    for (var i in nums) {
        if (i == 0 || nums[i] != nums[i-1]) {
            nums[l] = nums[i];
            l += 1;
        }
    }
    return l
};
