package leetcodego

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isEvenOddTree(root *TreeNode) bool {
    //bfs, TC O(n), SC O(n)
    nextLevelSlice := []*TreeNode{root}
    curLevel := 0
    for len(nextLevelSlice) != 0{
        curVal := -1
        tempSlice := []*TreeNode{}
        for _, node := range nextLevelSlice{
            if(node.Val%2 == curLevel%2){
                return false
            }
            if(curVal != -1 && ((curLevel%2 == 1 && curVal <= node.Val)||(curLevel%2 == 0 && curVal >= node.Val))){
                return false
            }
            curVal = node.Val
            if(node.Left != nil){
                tempSlice = append(tempSlice, node.Left)
            }
            if(node.Right != nil){
                tempSlice = append(tempSlice, node.Right)
            }
        }
        curLevel++
        nextLevelSlice = tempSlice
    }
    return true
}