class MapSum(object):

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.d = {}
        

    def insert(self, key, val):
        """
        :type key: str
        :type val: int
        :rtype: void
        """
        self.d[key] = val
        

    def sum(self, prefix):
        """
        :type prefix: str
        :rtype: int
        """
        sum = 0
        for k in self.d:
            if k.startswith(prefix):
                sum += self.d[k]
        return sum
        


# Your MapSum object will be instantiated and called as such:
# obj = MapSum()
# obj.insert(key,val)
# param_2 = obj.sum(prefix)