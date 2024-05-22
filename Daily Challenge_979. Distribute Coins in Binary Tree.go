package leetcodego

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 import(
	"math"
 )
 var moves int
 func countMoves(node *TreeNode) int{
	//DFS, TC O(n), SC O(height of tree)
	 if node == nil{
		 return 0
	 }
	 leftCount, rightCount := countMoves(node.Left), countMoves(node.Right)
	 moves += int(math.Abs(float64(leftCount))) + int(math.Abs(float64(rightCount)))
	 return node.Val + leftCount + rightCount - 1
 }
 func distributeCoins(root *TreeNode) int {
	 moves = 0
	 countMoves(root)
	 return moves
 }