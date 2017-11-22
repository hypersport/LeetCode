class Solution(object):
    def searchInsert(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        for i in range(len(nums)):
            if nums[i] > target:
                return 0
            elif nums[i] == target:
                return i
            elif i == len(nums)-1 or nums[i] < target and nums[i+1] > target:
                return i+1