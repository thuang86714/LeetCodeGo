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
 var diff2 int
 
 func findDiff(root *TreeNode) {
	 if root == nil {
		 return
	 }
 
	 findDiff(root.Left)
 
	 if prevNode != nil {
		 diff2 = min(diff2, root.Val-prevNode.Val)
	 }
	 prevNode = root
	 findDiff(root.Right)
 }
 func getMinimumDifference2(root *TreeNode) int {
	 diff2 = 100000
	 prevNode = nil
	 findDiff(root)
	 return diff2
 }