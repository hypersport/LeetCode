class Solution(object):
    def removeElement(self, nums, val):
        """
        :type nums: List[int]
        :type val: int
        :rtype: int
        """
        n = 0
        for i in range(len(nums)):
            if nums[i] != val:
                nums[i], nums[n] = nums[n], nums[i]
                n += 1
        return n
