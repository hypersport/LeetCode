class Solution(object):
    def maxSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        sum = n = nums[0]
        for i in range(1, len(nums)):
            if n > 0:
                n += nums[i]
            else:
                n = nums[i]
            if n > sum:
                sum = n
        return sum