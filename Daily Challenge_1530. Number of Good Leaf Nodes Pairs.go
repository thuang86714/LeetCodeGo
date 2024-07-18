package leetcodego

var ans int

func findPathDFS(root *TreeNode, distance int) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil{
		return []int{1}
	}

	leftSlice := findPathDFS(root.Left, distance)
	rightSlice := findPathDFS(root.Right, distance)

	for _, left := range leftSlice {
		for _, right := range rightSlice {
			if left+right <= distance {
				ans++
			}
		}
	}

	result := []int{}

	for _, left := range leftSlice {
		if left < distance {
			result = append(result, left+1)
		}
	}
	for _, right := range rightSlice {
		if right < distance {
			result = append(result, right+1)
		}
	}
	return result
}
func countPairs(root *TreeNode, distance int) int {
	ans = 0
	findPathDFS(root, distance)
	return ans
}

/*
func countPairs(root *TreeNode, distance int) int {
	ans := 0
	var findPathDFS func(*TreeNode, int) []int
	findPathDFS = func(root *TreeNode, distance int) []int {
		if root == nil {
			return []int{}
		}
		if root.Left == nil && root.Right == nil {
			return []int{1}
		}

		leftSlice := findPathDFS(root.Left, distance)
		rightSlice := findPathDFS(root.Right, distance)

		for _, left := range leftSlice {
			for _, right := range rightSlice {
				if left+right <= distance {
					ans++
				}
			}
		}

		result := []int{}

		for _, left := range leftSlice {
			if left < distance {
				result = append(result, left+1)
			}
		}
		for _, right := range rightSlice {
			if right < distance {
				result = append(result, right+1)
			}
		}
		return result
	}
	findPathDFS(root, distance)
	return ans
}
*/