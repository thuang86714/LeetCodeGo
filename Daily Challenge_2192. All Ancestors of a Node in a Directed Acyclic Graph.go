package leetcodego

import "sort"
//TC O(V*(V+E)*V^2logv), SC O(V^2 + E)
var directedGraphSlice [][]int

func buildGraph3(edges [][]int) {
	for _, edge := range edges {
		directedGraphSlice[edge[1]] = append(directedGraphSlice[edge[1]], edge[0])
	}
}

func getAncestors(n int, edges [][]int) [][]int {
	// Initialize slices
	directedGraphSlice = make([][]int, n)
	ansSlice := make([][]int, n)

	for idx := range directedGraphSlice {
		directedGraphSlice[idx] = []int{}
	}
	buildGraph3(edges)

	// Perform BFS for each node
	for idx := 0; idx < n; idx++ {
		visitedMap := make(map[int]bool)
		toVisitSlice := []int{idx}

		for len(toVisitSlice) != 0 {
			curNode := toVisitSlice[0]
			toVisitSlice = toVisitSlice[1:]

			for _, ancestor := range directedGraphSlice[curNode] {
				if !visitedMap[ancestor] {
					visitedMap[ancestor] = true
					ansSlice[idx] = append(ansSlice[idx], ancestor)
					toVisitSlice = append(toVisitSlice, ancestor)
				}
			}
		}

		sort.Ints(ansSlice[idx])
	}
	return ansSlice
}
