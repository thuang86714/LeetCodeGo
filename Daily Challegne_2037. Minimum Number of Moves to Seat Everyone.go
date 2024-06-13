package leetcodego

import(
	"sort"
	"math"
)

func minMovesToSeat(seats []int, students []int) int {
    //TC O(nlogn), SC O(1)-->counting sort could reach TC O(n), SC O(1)
    ans := 0
    sort.Ints(seats)
    sort.Ints(students)
    for idx := range seats{
        ans += int(math.Abs(float64(seats[idx] - students[idx])))
    }
    return ans
}