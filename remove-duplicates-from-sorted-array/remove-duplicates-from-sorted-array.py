class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        l = 0
        for i in range(len(nums)):
            if i == 0 or nums[i] != nums[i-1]:
                nums[l] = nums[i]
                l += 1
        return l
