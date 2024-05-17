package leetcodego
import (
	"container/heap"
)
type Point struct {
	safeness, x, y int
}

type MaxHeap []Point

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].safeness > h[j].safeness }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maximumSafenessFactor(grid [][]int) int {
	dir := [5]int{1, 0, -1, 0, 1}
	n := len(grid)
	safenessQueue := make([][2]int, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				safenessQueue = append(safenessQueue, [2]int{i, j})
			}
		}
	}

	for len(safenessQueue) > 0 {
		cur := safenessQueue[0]
		safenessQueue = safenessQueue[1:]
		for d := 0; d < 4; d++ {
			x, y := cur[0]+dir[d], cur[1]+dir[d+1]
			if x >= 0 && y >= 0 && x < n && y < n && grid[x][y] == 0 {
				grid[x][y] = grid[cur[0]][cur[1]] + 1
				safenessQueue = append(safenessQueue, [2]int{x, y})
			}
		}
	}

	dijSafenessMaxHeap := &MaxHeap{}
	heap.Init(dijSafenessMaxHeap)
	heap.Push(dijSafenessMaxHeap, Point{grid[0][0], 0, 0})

	for {
		top := heap.Pop(dijSafenessMaxHeap).(Point)
		sf, i, j := top.safeness, top.x, top.y
		if i == n-1 && j == n-1 {
			return sf - 1
		}
		for d := 0; d < 4; d++ {
			x, y := i+dir[d], j+dir[d+1]
			if x >= 0 && y >= 0 && x < n && y < n && grid[x][y] > 0 {
				heap.Push(dijSafenessMaxHeap, Point{min(sf, grid[x][y]), x, y})
				grid[x][y] *= -1
			}
		}
	}
}