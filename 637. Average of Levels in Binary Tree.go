package leetcodedynamicprogramming

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

 func averageOfLevels(root *TreeNode) []float64 {
    //BFS, credit to brianchang_tw, TC O(n), SC O(n)
    if root == nil{
        return []float64{}//handle edge case
    }

    curLevelQ := []*TreeNode{root}
    avgPerLevel := []float64{}
    for len(curLevelQ) != 0{
        nextLevelQ := []*TreeNode{}
        curLevelSum := 0.0
        for _, node := range curLevelQ{
            //update current level sum
            if node != nil{
                curLevelSum += float64(node.Val)
            }

            if node.Left != nil{
                nextLevelQ = append(nextLevelQ, node.Left)
            }

            if node.Right != nil{
                nextLevelQ = append(nextLevelQ, node.Right)
            }
        }

        curAvg := curLevelSum/float64(len(curLevelQ))
        avgPerLevel = append(avgPerLevel, curAvg)
        curLevelQ = nextLevelQ
    }
    return avgPerLevel
}