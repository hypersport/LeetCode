class Solution(object):
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        n = 0 
        for i in range(len(nums)):
            if nums[i]:
                nums[i],nums[n] = nums[n],nums[i]
                n += 1