/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 package leetcodego
 func diameterOfBinaryTree(root *TreeNode) int {
    res := 0
    var findDepth func(*TreeNode) int
    findDepth = func(node *TreeNode) int{
        if node == nil{
            return 0
        }
        leftHeight := findDepth(node.Left)
        rightHeight := findDepth(node.Right)
        res = max(res, 1 + leftHeight + rightHeight)
        return max(leftHeight, rightHeight) + 1
    }
    findDepth(root)
    return res - 1
}