class Solution:
    def findContentChildren(self, g: List[int], s: List[int]) -> int:
        g.sort()
        s.sort()
        result = 0
        g_len = len(g) - 1
        s_len = len(s) - 1
        while min(g_len, s_len) >= 0:
            if s[s_len] >= g[g_len]:
                result += 1
                s_len -= 1
            g_len -= 1
        return result
