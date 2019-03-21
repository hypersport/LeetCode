class Solution:
    def sortArrayByParity(self, A: List[int]) -> List[int]:
        left, right = 0, len(A)-1
        while left < right:
            if A[left] % 2 and not A[right] % 2:
                A[left], A[right] = A[right], A[left]
                left += 1
                right -= 1
            else:
                if not A[left] % 2:
                    left += 1
                if A[right] % 2:
                    right -= 1
        return A

# one line solution
# return [i for i in A if not i%2] + [i for i in A if i%2]
