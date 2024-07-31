package leetcodego

type Node struct {
    Val int
    Left *Node
    Right *Node
    Next *Node
}

 func connect(root *Node) *Node {
	//bfs, TC O(n), SC O(n)
    if root == nil{
        return root
    }
    curLevelSlice := []*Node{root}
    for len(curLevelSlice) != 0{
        curSize := len(curLevelSlice)
        tempSlice := []*Node{}
        for i := 0; i < curSize; i++{
            curNode := curLevelSlice[i]
            if i != curSize - 1{
                curNode.Next = curLevelSlice[i + 1]
            }
            if curNode.Left != nil{
                tempSlice = append(tempSlice, curNode.Left)
            }
            if curNode.Right != nil{
                tempSlice = append(tempSlice, curNode.Right)
            }
        }
        curLevelSlice = tempSlice
    }
    return root
}