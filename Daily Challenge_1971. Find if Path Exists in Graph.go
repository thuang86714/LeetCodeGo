package leetcodego

func buildGraph(edges [][]int) map[int][]int{
    graph := make(map[int][]int)
    for _, edge := range edges{
        graph[edge[0]] = append(graph[edge[0]], edge[1])
        graph[edge[1]] = append(graph[edge[1]], edge[0])
    }
    return graph
}

func validPath(n int, edges [][]int, source int, destination int) bool {
    //BFS, TC O(V+E), SC O(V^2)
    if source == destination{
        return true
    }

    visitedMap := make(map[int]bool)
    visitedMap[source] = true

    toVisitSlice := []int{}
    toVisitSlice = append(toVisitSlice, source)

    graph := buildGraph(edges)
    if len(graph[source]) == 0 || len(graph[destination]) == 0{
        return false
    }

    for len(toVisitSlice) != 0{
        curVertex := toVisitSlice[0]
        toVisitSlice = toVisitSlice[1:]
		
        for _, nextVertex := range graph[curVertex]{
            if nextVertex == destination{
                return true
            }

            if _, ok := visitedMap[nextVertex]; !ok{
                visitedMap[nextVertex] = true
                toVisitSlice = append(toVisitSlice, nextVertex)
            }
        }
    }
    return false
}