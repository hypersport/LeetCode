class Solution(object):
    def findUnsortedSubarray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        sorted_nums = sorted(nums)
        indexes = []
        if nums == sorted_nums:
            return 0
        for i in range(len(nums)):
            if nums[i] != sorted_nums[i]:
                indexes.append(i)
        return indexes[len(indexes) - 1] - indexes[0] + 1