package leetcodedynamicprogramming
import(
	"sort"
)
func sequentialDigits(low int, high int) []int {
    //credit to pratamandiko, TC O(1), SC O(1)
    ans := make([]int, 0)

    for i:= 1; i <= 9; i++{
        num := i
        for j := i + 1; j <= 9; j++{
            num = num*10 + j

            if num >= low && num <= high{
                ans = append(ans, num)
            }
        } 
    }
    sort.Ints(ans)
    return ans
}