class Solution(object):
    def countBinarySubstrings(self, s):
        """
        :type s: str
        :rtype: int
        """
        starti, lasti, sum = 0, 0, 0
        for i in range(1, len(s)):
            if s[i] != s[i-1]:
                sum += 1
                lasti = i - starti
                starti = i
            elif i - starti < lasti:
                sum += 1
        return sum