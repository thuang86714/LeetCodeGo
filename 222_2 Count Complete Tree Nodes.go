package leetcodego

func countNodes2(root *TreeNode) int {
	//TC O(logn)
	curL, curR := root, root
	leftHeight, rightHeight := 0, 0
	for curL != nil {
		curL = curL.Left
		leftHeight++
	}

	for curR != nil {
		curR = curR.Right
		rightHeight++
	}

	if leftHeight == rightHeight {
		return (1 << leftHeight) - 1
	}
	return 1 + countNodes2(root.Left) + countNodes2(root.Right)
}