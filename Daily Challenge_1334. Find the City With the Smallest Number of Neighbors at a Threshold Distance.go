package leetcodego

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
    //credit to lee215, Floyd Algorithm
    graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, n)
		for j := range graph[i] {
			graph[i][j] = 10001
		}
	}
    res, smallest := 0, n

    for _, edge := range edges{
        graph[edge[0]][edge[1]], graph[edge[1]][edge[0]] = edge[2], edge[2]
    }
    for i := 0; i < n; i++{
        graph[i][i] = 0
    }

    for k := 0; k < n; k++{
        for i := 0; i < n; i++{
            for j := 0; j < n; j++{
                graph[i][j] = min(graph[i][j], graph[i][k] + graph[k][j])
            }
        }
    }

    for i := 0 ;i < n; i++{
        count := 0
        for j := 0; j < n; j++{
            if graph[i][j] <= distanceThreshold{
                count++
            }
        }
        if count <= smallest{
            res = i
            smallest = count
        }
    }
    return res
}