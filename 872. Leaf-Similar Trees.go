/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 //partial credit to prodonik
func traverseAllLeaves(root *TreeNode, leaves *[]int) {//2nd arg is pass by reference
    //dfs, TC O(n), SC O(logn)
    if root == nil{
        return
    }
    if root.Left == nil && root.Right == nil{//if a node doesn't have any child, it's a leaf
        *leaves = append(*leaves, root.Val)
        return
    }
    traverseAllLeaves(root.Left, leaves)
    traverseAllLeaves(root.Right, leaves)
}
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    leafSlice1 := make([]int, 0)
    leafSlice2 := make([]int, 0)
    traverseAllLeaves(root1, &leafSlice1)//pass by reference
    traverseAllLeaves(root2, &leafSlice2)
    return reflect.DeepEqual(leafSlice1, leafSlice2)
}