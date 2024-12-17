package leetcodego

func xorQueries(arr []int, queries [][]int) []int {
    res := []int{}

    for idx := 1; idx < len(arr); idx++{
        arr[idx] ^= arr[idx - 1]
    }

    for _, q := range queries{
        if q[0] == 0{
            res = append(res, arr[q[1]])
        }else{
            res = append(res, arr[q[0] - 1] ^ arr[q[1]])
        }
    }
    return res
}