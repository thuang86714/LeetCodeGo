/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 package leetcodego

 func findSum(root *TreeNode, sum int) int {
    //DFS TC O(n), SC O(logn)
    ans := 0
    sum = sum*10 + root.Val
    if root.Left == nil && root.Right == nil{
        return sum
    }
    if root.Left != nil{
        ans += findSum(root.Left, sum)
    }
    if root.Right != nil{
        ans += findSum(root.Right, sum)
    }
    return ans
}
func sumNumbers(root *TreeNode) int {
    return findSum(root, 0)
}