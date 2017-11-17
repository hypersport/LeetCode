class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype List[int]
        """
        d = {}
        if len(nums) < 1:
            return false
        for i in range(len(nums)):
            if nums[i] in d:
                return [d[nums[i]], i]
            d[target-nums[i]] = i
