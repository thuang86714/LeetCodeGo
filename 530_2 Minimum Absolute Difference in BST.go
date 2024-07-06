package leetcodego
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var prevNode *TreeNode
 var diff int
 
 func findDiff(root *TreeNode) {
	 if root == nil {
		 return
	 }
 
	 findDiff(root.Left)
 
	 if prevNode != nil {
		 diff = min(diff, root.Val-prevNode.Val)
	 }
	 prevNode = root
	 findDiff(root.Right)
 }
 func getMinimumDifference2(root *TreeNode) int {
	 diff = 100000
	 prevNode = nil
	 findDiff(root)
	 return diff
 }