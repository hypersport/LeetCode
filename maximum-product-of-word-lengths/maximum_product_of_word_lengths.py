class Solution:
    def maxProduct(self, words: List[str]) -> int:
        result = 0
        for i in range(len(words)):
            for j in range(i+1, len(words)):
                if not set(words[i]) & set(words[j]):
                    result = max(result, len(words[i]) * len(words[j]))
        return result
