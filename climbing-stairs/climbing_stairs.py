class Solution(object):
    def climbStairs(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n <= 0:
            return 0
        s = [0,1]
        for i in range(n):
            s[0],s[1] = s[1],s[0]+s[1]
        return s[1]