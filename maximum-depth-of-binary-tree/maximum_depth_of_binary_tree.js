/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */
var maxDepth = function(root) {
    if (!root) {
        return 0;
    }
    var left = maxDepth(root.left);
    var right = maxDepth(root.right);
    if (!left || !right) {
        return left + right + 1;
    } else if (left >= right) {
        return left + 1;
    } else {
        return right + 1;
    }
};