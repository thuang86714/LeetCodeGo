package leetcodego

func minimizeArrayValue(nums []int) int {
    //credit to lee215, TC O(n), SC O(1)
    /*
    The ceil average with sum of i + 1 number:
    ceil(double(sum) / (i + 1))

    We can also do integer division:
    (sum + i) / (i + 1)
    */
    sum, res := 0, 0
    for idx, val := range nums{
        sum += val
        res = max(res, (sum + idx)/(idx + 1))
    }
    return res
}