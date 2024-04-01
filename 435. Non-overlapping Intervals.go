package leetcodego

import(
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int)bool{//do 2D-slice sorting by their end elements
        return intervals[i][1] < intervals[j][1]
    })

    count, curBack := 0, intervals[0][1]
    for i := 1; i < len(intervals); i++{
        if intervals[i][0] < curBack{
            count++
        }else{
            curBack = intervals[i][1]
        }
    }
    return count
}