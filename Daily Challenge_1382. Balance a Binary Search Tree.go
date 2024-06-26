package leetcodego

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var treeNodeSlice []*TreeNode
 func collectAllNodes(node *TreeNode) {
	 if node == nil{
		 return
	 }
	 collectAllNodes(node.Left)
	 treeNodeSlice = append(treeNodeSlice, node)
	 collectAllNodes(node.Right)
 }
 func buildTree(start, end int) *TreeNode {
	 if start > end {
		 return nil
	 }
	 mid := start + (end-start)/2
	 node := treeNodeSlice[mid]
	 node.Left = buildTree(start, mid-1)
	 node.Right = buildTree(mid+1, end)
	 return node
 }
 func balanceBST(root *TreeNode) *TreeNode {
	 treeNodeSlice = []*TreeNode{}
	 //TC O(n), SC O(n), find the mid point and go to both ends to build balance tree
	 collectAllNodes(root)
	 totalNodeCount := len(treeNodeSlice)
	 return buildTree(0, totalNodeCount - 1)
 }