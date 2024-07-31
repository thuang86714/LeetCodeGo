/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

 

Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
Example 2:

Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping
*/
package leetcodego
import "sort"
func merge4(intervals [][]int) [][]int {
	//credit to brianchiang_tw
    result := [][]int{}
    sort.Slice(intervals, func(a, b int) bool{
        return (intervals[a][0] < intervals[b][0]) || ((intervals[a][0]==intervals[b][0])&&(intervals[a][1]<intervals[b][1]))
    })
    for _, cur:= range intervals{
        if len(result)==0 || result[len(result)-1][1] < cur[0]{
            result = append(result, cur)
        }else{
            // has overlapping
            // merge with previous interval
            result[len(result)-1][1] = max(result[len(result)-1][1], cur[1])
        }
    }
    return result
}

func max(x, y int) int{
    if x >= y{
        return x
    }else{
        return y
    }

}