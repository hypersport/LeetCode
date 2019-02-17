class Solution:
    def licenseKeyFormatting(self, S: 'str', K: 'int') -> 'str':
        a = ''.join(S.split('-'))[::-1]
        res = ''
        i = 0
        for c in a:
            if i != K:
                res += c.upper()
                i += 1
            else:
                res += '-' + c.upper()
                i = 1
        return res[::-1]
