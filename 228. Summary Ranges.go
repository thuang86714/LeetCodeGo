package leetcodego
import(
	"strconv"
)
func summaryRanges(nums []int) []string {
    if len(nums) == 0{
        return []string{}
    }
    res, curBack, curFront := []string{}, nums[0], nums[0]
    for _, val := range nums{
        if curBack + 1 < val{
            if curFront == curBack{
                res = append(res, strconv.Itoa(curFront))
            }else{
                res = append(res, strconv.Itoa(curFront) + "->" + strconv.Itoa(curBack))
            }
            curFront = val
        }
        curBack = val
    }

    if curFront == curBack{
        res = append(res, strconv.Itoa(curFront))
    }else{
        res = append(res, strconv.Itoa(curFront) + "->" + strconv.Itoa(curBack))
    }

    return res
}