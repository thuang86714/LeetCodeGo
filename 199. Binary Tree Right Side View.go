/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 package leetcodego
 func rightSideView(root *TreeNode) []int {
    //BFS, TC O(n), SC O(n)
    if root == nil{
        return []int{}
    }
    nextLevelSlice := []*TreeNode{root}
    ans := []int{}
    for len(nextLevelSlice) != 0{
        tempSlice := []*TreeNode{}
        var rightNode int
        for _, node := range nextLevelSlice{
            rightNode = node.Val//the right-most node would eventually be stored in rightNode
            if node.Left != nil{
                tempSlice = append(tempSlice, node.Left)
            }
            if node.Right != nil{
                tempSlice = append(tempSlice, node.Right)
            }
        }
        ans = append(ans, rightNode)
        nextLevelSlice = tempSlice
    }
    return ans
}