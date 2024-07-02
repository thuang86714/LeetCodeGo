package leetcodego
import (
    "container/heap"
)
type IntHeap []int
func (h IntHeap) Len() int{ return len(h)}
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j]}
func (h IntHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}

func (h *IntHeap) Push(x interface{}){
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func findKthLargest(nums []int, k int) int {
    //credit to vanAmsen, in golang, there's only interface for the minHeap. Need to implement it ourselves
    minHeap := IntHeap(nums[:k])//initializes minHeap with the first k elements of nums. minHeap is an instance of IntHeap, which implements the heap.Interface
    heap.Init(&minHeap)//called to transform the minHeap slice into a valid heap, according to the Less method defined by IntHeap (in your case, it ensures that IntHeap functions as a min heap).

    for _, num := range nums[k:]{
        if num > minHeap[0]{
            heap.Pop(&minHeap)
            heap.Push(&minHeap, num)
        }
    }
    return minHeap[0]
}