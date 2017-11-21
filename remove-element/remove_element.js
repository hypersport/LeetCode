/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
var removeElement = function(nums, val) {
    var n = 0;
    var tmp;
    for(var i = 0; i < nums.length; i ++) {
        if (nums[i] != val) {
            tmp = nums[i];
            nums[i] = nums[n];
            nums[n] = tmp
            n += 1
        }
    }
    return n;
};
