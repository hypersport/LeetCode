class Solution:
    def isValid(self, s: str) -> bool:
        parentheses = []

        for c in s:
            if c == '(' or c == '[' or c == '{':
                parentheses.append(c)
            else:
                if parentheses:
                    currentBracket = parentheses.pop()
                    if c == ')' and currentBracket != '(':
                        return False
                    elif c == ']' and currentBracket != '[':
                        return False
                    elif c == '}' and currentBracket != '{':
                        return False
                else:
                    if c == ')' or c == ']' or c == '}':
                        return False

        if parentheses:
            return False
        return True
