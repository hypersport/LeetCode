/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    var d = {};
    var t;
    for (var i in nums) {
        t = target - nums[i];
        if (t in d) {
            return new Array(parseInt(d[t]), parseInt(i));
        } else {
            d[nums[i]] = i;
        }
    }
    return null;
};