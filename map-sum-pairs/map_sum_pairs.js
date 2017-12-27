/**
 * Initialize your data structure here.
 */
var MapSum = function() {
    this.d = new Array();
};

/** 
 * @param {string} key 
 * @param {number} val
 * @return {void}
 */
MapSum.prototype.insert = function(key, val) {
    this.d[key] = val;
};

/** 
 * @param {string} prefix
 * @return {number}
 */
MapSum.prototype.sum = function(prefix) {
    var sum = 0;
    var sLen = prefix.length;
    for (var key in this.d) {
        if (key.slice(0, sLen) == prefix) {
            sum += this.d[key];
        }
    }
    return sum
};

/** 
 * Your MapSum object will be instantiated and called as such:
 * var obj = Object.create(MapSum).createNew()
 * obj.insert(key,val)
 * var param_2 = obj.sum(prefix)
 */