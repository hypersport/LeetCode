class Solution(object):
    def isSubsequence(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        index = 0
        for i in range(len(s)):
            index = t.find(s[i], index)
            if index < 0:
                return False
            index += 1
        return True