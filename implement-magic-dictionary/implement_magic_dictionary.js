/**
 * Initialize your data structure here.
 */
var MagicDictionary = function() {
    this.rdict = null;
};

/**
 * Build a dictionary through a list of words
 * @param {string[]} dict
 * @return {void}
 */
MagicDictionary.prototype.buildDict = function(dict) {
    this.rdict = dict;
};

/**
 * Returns if there is any word in the trie that equals to the given word after modifying exactly one character
 * @param {string} word
 * @return {boolean}
 */
MagicDictionary.prototype.search = function(word) {
    for (var index in this.rdict) {
        var n = 0;
        if (word.length == this.rdict[index].length) {
            for (var i = 0; i < word.length; i ++) {
                if (word.charAt(i) != this.rdict[index].charAt(i)) {
                    n ++;
                }
            }
            if (n == 1) {
                return true;
            }
        }
    }
    return false;
};

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * var obj = Object.create(MagicDictionary).createNew()
 * obj.buildDict(dict)
 * var param_2 = obj.search(word)
 */
