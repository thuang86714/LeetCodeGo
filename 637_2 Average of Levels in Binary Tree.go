package leetcodego

func averageOfLevels2(root *TreeNode) []float64 {
	//BFS, TC O(n)
	toVisitSlice := []*TreeNode{root}
	ansSlice := []float64{}

	for len(toVisitSlice) != 0 {
		curSum, curCount := 0.0, len(toVisitSlice)
		curSize := float64(curCount)
		for curCount > 0 {
			curNode := toVisitSlice[0]
			toVisitSlice = toVisitSlice[1:]
			curSum += float64(curNode.Val)
			if curNode.Left != nil {
				toVisitSlice = append(toVisitSlice, curNode.Left)
			}
			if curNode.Right != nil {
				toVisitSlice = append(toVisitSlice, curNode.Right)
			}
			curCount--
		}
		ansSlice = append(ansSlice, curSum/curSize)
	}
	return ansSlice
}