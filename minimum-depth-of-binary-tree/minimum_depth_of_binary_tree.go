/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := minDepth(root.Left)
    right := minDepth(root.Right)
    if left == 0 || right == 0 {
        return left + right + 1
    } else if left >= right {
        return right + 1
    } else {
        return left + 1
    }
}