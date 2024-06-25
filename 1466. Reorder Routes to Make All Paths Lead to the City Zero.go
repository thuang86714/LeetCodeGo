package leetcodego

var directedGraph map[int]map[int]bool
var unDirectedGraph map[int][]int
var visitedMap map[int]bool

func buildGraph2(connections [][]int) {
	directedGraph = make(map[int]map[int]bool) // Initialize the directed graph
	unDirectedGraph = make(map[int][]int)      // Initialize the undirected graph
	for _, connection := range connections {
		// For directedGraph, make sure the map for connection[1] is initialized
		if directedGraph[connection[1]] == nil {
			directedGraph[connection[1]] = make(map[int]bool)
		}
		// For directedGraph, also ensure the map for connection[0] is initialized for checking in minReorder
		if directedGraph[connection[0]] == nil {
			directedGraph[connection[0]] = make(map[int]bool)
		}
		unDirectedGraph[connection[0]] = append(unDirectedGraph[connection[0]], connection[1])
		unDirectedGraph[connection[1]] = append(unDirectedGraph[connection[1]], connection[0])
		directedGraph[connection[1]][connection[0]] = true
	}
}
func minReorder(n int, connections [][]int) int {
	//BFS TC O(n + e), SC O(n + e)
	buildGraph2(connections)
	visitedMap = make(map[int]bool) // Initialize the visited map
	toVisitSlice := []int{}
	reorderCount := 0
	toVisitSlice = append(toVisitSlice, 0)
	for len(toVisitSlice) != 0 {
		curLevelSize := len(toVisitSlice)
		for i := 0; i < curLevelSize; i++ {
			curNode := toVisitSlice[0]
			toVisitSlice = toVisitSlice[1:]
			visitedMap[curNode] = true
			for _, nextNode := range unDirectedGraph[curNode] {
				if visitedMap[nextNode] == true {
					continue
				}
				if _, exist := directedGraph[curNode][nextNode]; !exist {
					reorderCount++
				}
				toVisitSlice = append(toVisitSlice, nextNode)
			}
		}
	}
	return reorderCount
}