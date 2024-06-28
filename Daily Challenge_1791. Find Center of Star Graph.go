package leetcodego

func findCenter(edges [][]int) int {
	//find the node with highest degree, TC O(n), SC O(n)
	n, ans, curCount := len(edges)+1, 0, 0
	degreeCountSlice := make([]int, n+1)

	for _, edge := range edges {
		degreeCountSlice[edge[0]]++
		degreeCountSlice[edge[1]]++
	}

	for idx, degree := range degreeCountSlice {
		if degree > curCount {
			curCount = degree
			ans = idx
		}
	}

	return ans
}

/*
func findCenter(edges [][]int) int {
	// credit to votrubac, TC O(1), SC O(1)
	// Basically, a center node must appear in every edge.
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}
	return edges[0][1]
}
*/