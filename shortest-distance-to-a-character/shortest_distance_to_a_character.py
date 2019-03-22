class Solution:
    def shortestToChar(self, S: str, C: str) -> List[int]:
        indexes = [i for i in range(len(S)) if S[i] == C]
        result = [float('inf') for c in S]
        index, i, stop = 0, 0, 0

        for index in range(1, len(indexes)):
            for i in range(stop, len(S)):
                if i == indexes[index]:
                    result[i] = 0
                    stop = i+1
                    break
                elif i <= indexes[index-1]:
                    result[i] = indexes[index-1]-i
                elif i > indexes[index-1] and i < indexes[index]:
                    result[i] = min(i-indexes[index-1], indexes[index]-i)

        while len(S) > i:
            result[i] = max(i-indexes[index], indexes[index]-i)
            i += 1

        return result
