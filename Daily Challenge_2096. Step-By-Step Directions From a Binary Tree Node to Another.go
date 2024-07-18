package leetcodego

func findPathStr(node *TreeNode, value int, curByteSlice *[]byte) bool {
	if node.Val == value {
		return true
	}

	if node.Left != nil && findPathStr(node.Left, value, curByteSlice) {
		*curByteSlice = append(*curByteSlice, 'L')
	} else if node.Right != nil && findPathStr(node.Right, value, curByteSlice) {
		*curByteSlice = append(*curByteSlice, 'R')
	}

	return len(*curByteSlice) > 0
}
func getDirections(root *TreeNode, startValue int, destValue int) string {
	var destByteSlice []byte
	var startByteSlice []byte
	findPathStr(root, destValue, &destByteSlice)
	findPathStr(root, startValue, &startByteSlice)
	for len(destByteSlice) != 0 && len(startByteSlice) != 0 && destByteSlice[len(destByteSlice)-1] == startByteSlice[len(startByteSlice)-1] {
		destByteSlice = destByteSlice[:len(destByteSlice)-1]
		startByteSlice = startByteSlice[:len(startByteSlice)-1]
	}

	temp := make([]byte, len(startByteSlice))
	for idx := range temp {
		temp[idx] = 'U'
	}

	for i, j := 0, len(destByteSlice)-1; i < j; {
		destByteSlice[i], destByteSlice[j] = destByteSlice[j], destByteSlice[i]
		i++
		j--
	}

	temp = append(temp, destByteSlice...)
	return string(temp)
}