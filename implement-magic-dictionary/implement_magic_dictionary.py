class MagicDictionary(object):

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.rdict = {}
        

    def buildDict(self, dict):
        """
        Build a dictionary through a list of words
        :type dict: List[str]
        :rtype: void
        """
        i = 0
        for word in dict:
            self.rdict[i] = word
            i += 1
        

    def search(self, word):
        """
        Returns if there is any word in the trie that equals to the given word after modifying exactly one character
        :type word: str
        :rtype: bool
        """
        for oword in self.rdict.values():
            n = 0
            if len(oword) == len(word):
                for i in range(len(word)):
                    if oword[i] != word[i]:
                        n += 1
                if n == 1:
                    return True
        return False
        


# Your MagicDictionary object will be instantiated and called as such:
# obj = MagicDictionary()
# obj.buildDict(dict)
# param_2 = obj.search(word)