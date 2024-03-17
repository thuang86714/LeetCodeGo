/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func maxLevelSum(root *TreeNode) int {
    //BFS TC O(n), SC O(n)
    curMax := math.MinInt
    maxLevel, curLevel := 1, 1
    nodeSlice := []*TreeNode{root}
    for len(nodeSlice) > 0{
        curSize := len(nodeSlice)
        curSum := 0
        tempSlice := []*TreeNode{}
        for i := 0; i < curSize; i++{
            node := nodeSlice[i]
            if node.Left != nil{
                tempSlice = append(tempSlice, node.Left)
            }
            if node.Right != nil{
                tempSlice = append(tempSlice, node.Right)
            }
            curSum += node.Val
        }
        if curSum > curMax{
            maxLevel = curLevel
            curMax = curSum
        }
        curLevel++
        nodeSlice = tempSlice
    }
    return maxLevel
}