class Solution:
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        if k == 1 or len(nums) == 1:
            return max(nums)/1.0
        maxKSum = currentSum = sum(nums[:k])
        for i in range(len(nums)-k):
            currentSum = currentSum - nums[i] + nums[i+k]
            maxKSum = max(currentSum, maxKSum)
        return maxKSum/k*1.0
