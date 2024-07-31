/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 package leetcodego
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//DFS, TC O(n), SC O(logn)
	if root == nil{
		return root
	}else if root == p{
		return p
	}else if root == q{
		return q
	}

   left := lowestCommonAncestor(root.Left, p, q)
   right := lowestCommonAncestor(root.Right, p, q)

   if left == nil && right == nil{
	   return nil
   }else if left == nil{
	   return right
   }else if right == nil{
	   return left
   }else{
	   return root
   }
}