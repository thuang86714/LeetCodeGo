
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrder(root *TreeNode) [][]int {
    //BFS, TC O(n), SC O(n)
    if root == nil{
        return [][]int{}
    }
    ans := [][]int{}
    nextLevelSlice := []*TreeNode{root}
    for len(nextLevelSlice) != 0{
        thisLevelSlice := []int{}
        tempSlice := []*TreeNode{}
        curLen := len(nextLevelSlice)
        for i := 0; i < curLen; i++{
            curNode := nextLevelSlice[i]
            thisLevelSlice = append(thisLevelSlice, curNode.Val)
            if curNode.Left != nil{
                tempSlice = append(tempSlice, curNode.Left)
            }
            if curNode.Right != nil{
                tempSlice = append(tempSlice, curNode.Right)
            }
        }
        ans = append(ans, thisLevelSlice)
    
        nextLevelSlice = tempSlice
    }
    return ans
}