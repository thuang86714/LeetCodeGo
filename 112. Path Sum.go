/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 package leetcodego

 func hasPathSum(root *TreeNode, targetSum int) bool {
    //DFS, TC O(n), SC O(logn)
    if root == nil{
        return false
    }
    targetSum -= root.Val
    if targetSum == 0 && root.Left == nil && root.Right == nil{
        return true
    }
    return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}