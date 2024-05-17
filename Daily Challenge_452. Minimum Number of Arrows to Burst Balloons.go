package leetcodego
import(
    "sort"
)
func findMinArrowShots(points [][]int) int {
    //TC O(nlogn), SC O(1)
    arrow := 1
    sort.Slice(points, func(i, j int)bool{
        return points[i][0] < points[j][0]
    })
    arrowAt := points[0][1]
    for i := 1; i < len(points); i++{
        if arrowAt < points[i][0]{
            arrowAt = points[i][1]
            arrow++
        }else{
            arrowAt = min(arrowAt, points[i][1])
        }
    }
    return arrow
}